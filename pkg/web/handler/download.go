package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/enum"
	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/middleware"
	util2 "dumpapp_server/pkg/util"
	"github.com/go-playground/validator/v10"
	pkgErr "github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/volatiletech/null/v8"
)

type DownloadHandler struct {
	accountDAO                 dao.AccountDAO
	ipaDAO                     dao.IpaDAO
	ipaVersionDAO              dao.IpaVersionDAO
	memberDownloadIpaRecordDAO dao.MemberDownloadIpaRecordDAO
	cribberDAO                 dao.CribberDAO

	memberDownloadNumberCtl controller.MemberPayCountController
	tencentCtl              controller.TencentController
	lingshulianCtl          controller.LingshulianController
	ipaVersionCtl           controller.IpaVersionController
}

func NewDownloadHandler() *DownloadHandler {
	return &DownloadHandler{
		accountDAO:                 impl.DefaultAccountDAO,
		ipaDAO:                     impl.DefaultIpaDAO,
		ipaVersionDAO:              impl.DefaultIpaVersionDAO,
		memberDownloadIpaRecordDAO: impl.DefaultMemberDownloadIpaRecordDAO,
		cribberDAO:                 impl.DefaultCribberDAO,

		memberDownloadNumberCtl: impl2.DefaultMemberPayCountController,
		tencentCtl:              impl2.DefaultTencentController,
		lingshulianCtl:          impl2.DefaultLingshulianController,
		ipaVersionCtl:           impl2.DefaultIpaVersionController,
	}
}

type checkCanDownloadArgs struct {
	Version string `form:"version" validate:"required"`
	IpaType string `form:"ipa_type"`
}

func (args *checkCanDownloadArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if args.IpaType == "" {
		args.IpaType = enum.IpaTypeNormal.String()
	}
	return nil
}

func (h *DownloadHandler) CheckCanDownload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ipaID := cast.ToInt64(util.URLParam(r, "ipa_id"))

	args := checkCanDownloadArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	ipaType, err := enum.IpaTypeString(args.IpaType)
	util.PanicIf(err)

	loginID := middleware.MustGetMemberID(ctx)

	ipaVersions, err := h.ipaVersionDAO.GetByIpaIDAndIpaTypeAndVersion(ctx, ipaID, ipaType, args.Version)
	util.PanicIf(err)

	if len(ipaVersions) == 0 {
		util.RenderJSON(w, map[string]interface{}{
			"can_download": false,
			"message":      "此 ipa 没有检索到对应的版本",
		})
		return
	}

	/// 判断之前是否下载过
	dn, err := h.memberDownloadIpaRecordDAO.GetByMemberIDIpaIDIpaTypeVersion(ctx, loginID, null.Int64From(ipaID), null.StringFrom(ipaType.String()), null.StringFrom(args.Version))
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}

	message := ""
	if dn == nil {
		message = "花费 9 个 D 币进行下载（有任何异常问题，可联系客服返还 D 币）"
	}

	resJSON := map[string]interface{}{
		"can_download": true,
		"message":      message,
	}
	util.RenderJSON(w, resJSON)
}

type getDownloadURLArgs struct {
	Version string `form:"version" validate:"required"`
	IpaType string `form:"ipa_type"`
}

func (args *getDownloadURLArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if args.IpaType == "" {
		args.IpaType = enum.IpaTypeNormal.String()
	}
	return nil
}

func (h *DownloadHandler) GetDownloadURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ipaID := cast.ToInt64(util.URLParam(r, "ipa_id"))

	args := getDownloadURLArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	ipaType, err := enum.IpaTypeString(args.IpaType)
	util.PanicIf(err)

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

	dn, err := h.memberDownloadIpaRecordDAO.GetByMemberIDIpaIDIpaTypeVersion(ctx, loginID, null.Int64From(ipaID), null.StringFrom(ipaType.String()), null.StringFrom(args.Version))
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}

	/// 如果之前没有下载过, 则需要扣除 9 个积分
	if dn == nil {
		/// 事物
		txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
		defer clients.MustClearMySQLTransaction(ctx, txn)
		ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)
		/// 添加下载记录
		recordID := util2.MustGenerateID(ctx)
		/// 消费 9 个积分
		util.PanicIf(h.memberDownloadNumberCtl.DeductPayCount(ctx, loginID, constant.IpaPrice, enum.MemberPayCountStatusUsed, enum.MemberPayCountUseIpa, datatype.MemberPayCountRecordBizExt{
			ObjectID:   recordID,
			ObjectType: datatype.MemberPayCountRecordBizExtObjectTypeDownloadIpaRecord,
		}))
		util.PanicIf(h.memberDownloadIpaRecordDAO.Insert(ctx, &models.MemberDownloadIpaRecord{
			ID:       recordID,
			MemberID: loginID,
			Status:   "used",
			IpaID:    null.Int64From(ipaID),
			IpaType:  null.StringFrom(ipaType.String()),
			Version:  null.StringFrom(args.Version),
		}))
		clients.MustCommit(ctx, txn)
		ctx = util.ResetCtxKey(ctx, constant.TransactionKeyTxn)
	}

	ipaVersions, err := h.ipaVersionDAO.GetByIpaIDAndIpaTypeAndVersion(ctx, ipaID, ipaType, args.Version)
	util.PanicIf(err)
	if len(ipaVersions) == 0 {
		util.RenderJSON(w, map[string]interface{}{
			"can_download": false,
		})
		return
	}

	version := ipaVersions[0]

	openURL, err := h.ipaVersionCtl.GetDownloadURL(ctx, version.ID, loginID)
	util.PanicIf(err)

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
