package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/middleware"
	util2 "dumpapp_server/pkg/util"
	"github.com/go-playground/validator/v10"
	pkgErr "github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/volatiletech/null/v8"
)

type DownloadHandler struct {
	accountDAO              dao.AccountDAO
	ipaDAO                  dao.IpaDAO
	ipaVersionDAO           dao.IpaVersionDAO
	memberDownloadNumberDAO dao.MemberDownloadNumberDAO
	cribberDAO              dao.CribberDAO

	memberDownloadNumberCtl controller.MemberDownloadController
	tencentCtl              controller.TencentController
}

func NewDownloadHandler() *DownloadHandler {
	return &DownloadHandler{
		accountDAO:              impl.DefaultAccountDAO,
		ipaDAO:                  impl.DefaultIpaDAO,
		ipaVersionDAO:           impl.DefaultIpaVersionDAO,
		memberDownloadNumberDAO: impl.DefaultMemberDownloadNumberDAO,
		cribberDAO:              impl.DefaultCribberDAO,

		memberDownloadNumberCtl: impl2.DefaultMemberDownloadController,
		tencentCtl:              impl2.DefaultTencentController,
	}
}

type checkCanDownloadArgs struct {
	Version string `form:"version" validate:"required"`
}

func (args *checkCanDownloadArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *DownloadHandler) CheckCanDownload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ipaID := cast.ToInt64(util.URLParam(r, "ipa_id"))

	args := checkCanDownloadArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	loginID := middleware.MustGetMemberID(ctx)

	ipaVersion, err := h.ipaVersionDAO.GetByIpaIDVersion(ctx, ipaID, args.Version)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}
	if ipaVersion == nil {
		util.RenderJSON(w, map[string]interface{}{
			"can_download": false,
		})
		return
	}

	dn, err := h.memberDownloadNumberDAO.GetByMemberIDIpaIDVersion(ctx, loginID, null.Int64From(ipaID), null.StringFrom(args.Version))
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}

	resJSON := map[string]interface{}{
		"can_download": dn != nil,
	}
	util.RenderJSON(w, resJSON)
}

type getDownloadURLArgs struct {
	Version string `form:"version" validate:"required"`
}

func (args *getDownloadURLArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *DownloadHandler) GetDownloadURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ipaID := cast.ToInt64(util.URLParam(r, "ipa_id"))

	args := getDownloadURLArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	loginID := middleware.MustGetMemberID(ctx)

	/// 以下是一套反作弊的机制
	isBlackList, err := h.cribberDAO.GetBlacklistByMemberID(ctx, loginID)
	util.PanicIf(err)
	if isBlackList {
		panic(errors.ErrMemberBlacklist)
	}

	remoteIP := cast.ToString(ctx.Value(constant.RemoteIP))
	incrCount, err := h.cribberDAO.GetRemoteIPIncrCount(ctx, loginID, remoteIP)
	util.PanicIf(err)
	if incrCount > 10 {
		/// 加入黑名单
		util.PanicIf(h.cribberDAO.SetMemberIDToBlacklist(ctx, loginID))
		h.sendBlacklistWeiXinBot(ctx, loginID, ipaID, remoteIP)
		panic(errors.ErrMemberBlacklist)
	}

	/// 操作数 +1
	util.PanicIf(h.cribberDAO.IncrMemberRemoteIP(ctx, loginID, remoteIP))

	dn, err := h.memberDownloadNumberDAO.GetByMemberIDIpaIDVersion(ctx, loginID, null.Int64From(ipaID), null.StringFrom(args.Version))
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}

	/// 如果之前没有下载过, 则需要扣除一次下载次数
	if dn == nil {
		dn, err := h.memberDownloadNumberCtl.GetDownloadNumber(ctx, loginID)
		util.PanicIf(err)
		dn.Status = enum.MemberDownloadNumberStatusUsed
		dn.IpaID = null.Int64From(ipaID)
		dn.Version = null.StringFrom(args.Version)
		util.PanicIf(h.memberDownloadNumberDAO.Update(ctx, dn))
	}

	ipaVersion, err := h.ipaVersionDAO.GetByIpaIDVersion(ctx, ipaID, args.Version)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}
	if ipaVersion == nil {
		util.RenderJSON(w, map[string]interface{}{
			"can_download": false,
		})
		return
	}

	openURL, err := h.tencentCtl.GetSignatureURL(ctx, ipaVersion.TokenPath)
	util.PanicIf(err)
	openURL = fmt.Sprintf("%s&member_id=%d", openURL, loginID)
	resJSON := map[string]interface{}{
		"open_url": openURL,
	}

	util.RenderJSON(w, resJSON)
}

func (h *DownloadHandler) sendBlacklistWeiXinBot(ctx context.Context, loginID, ipaID int64, ip string) {
	account, err := h.accountDAO.Get(ctx, loginID)
	util.PanicIf(err)

	ipa, err := h.ipaDAO.Get(ctx, ipaID)
	util.PanicIf(err)

	ipStr := fmt.Sprintf("ip 地址：<font color=\"comment\">%s</font>\n", ip)
	idStr := fmt.Sprintf("ID：<font color=\"comment\">%d</font>\n", account.ID)
	email := fmt.Sprintf("邮箱：<font color=\"comment\">%s</font>\n", account.Email)
	number := fmt.Sprintf("手机号：：<font color=\"comment\">%s</font>\n", account.Phone)
	createdAtStr := fmt.Sprintf("注册时间：<font color=\"comment\">%s</font>\n", account.CreatedAt.Format("2006-01-02 15:04:05"))
	ipaIDStr := fmt.Sprintf("Ipa ID：<font color=\"comment\">%d</font>\n", ipa.ID)
	ipaName := fmt.Sprintf("Ipa 名称：<font color=\"comment\">%s</font>\n", ipa.Name)
	timeStr := fmt.Sprintf("发送时间：<font color=\"comment\">%s</font>\n", time.Now().Format("2006-01-02 15:04:05"))
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": "<font color=\"warning\">危险: 发现恶意刷流量用户</font>\n>" +
				ipStr + idStr + email + number + createdAtStr + ipaIDStr + ipaName + timeStr,
		},
	}
	util2.SendWeiXinBot(ctx, config.DumpConfig.AppConfig.TencentGroupKey, data, []string{"@all"})
}
