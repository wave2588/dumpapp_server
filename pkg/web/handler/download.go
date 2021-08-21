package handler

import (
	errors2 "dumpapp_server/pkg/common/errors"
	"fmt"
	pkgErr "github.com/pkg/errors"
	"net/http"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"github.com/volatiletech/null/v8"
)

type DownloadHandler struct {
	ipaDAO                  dao.IpaDAO
	ipaVersionDAO           dao.IpaVersionDAO
	memberDownloadNumberDAO dao.MemberDownloadNumberDAO

	memberDownloadNumberCtl controller.MemberDownloadController
	tencentCtl              controller.TencentController
}

func NewDownloadHandler() *DownloadHandler {
	return &DownloadHandler{
		ipaDAO:                  impl.DefaultIpaDAO,
		ipaVersionDAO:           impl.DefaultIpaVersionDAO,
		memberDownloadNumberDAO: impl.DefaultMemberDownloadNumberDAO,

		memberDownloadNumberCtl: impl2.DefaultMemberDownloadController,
		tencentCtl:              impl2.DefaultTencentController,
	}
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
	util.PanicIf(err)

	openURL, err := h.tencentCtl.GetSignatureURL(ctx, ipaVersion.TokenPath)
	util.PanicIf(err)
	resJSON := map[string]interface{}{
		"open_url": openURL,
	}

	util.RenderJSON(w, resJSON)
}
