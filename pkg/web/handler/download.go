package handler

import (
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/middleware"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"net/http"
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

	dn, err := h.memberDownloadNumberCtl.GetDownloadNumber(ctx, loginID)
	util.PanicIf(err)

	dn.Status = enum.MemberDownloadNumberStatusUsed
	util.PanicIf(h.memberDownloadNumberDAO.Update(ctx, dn))

	ipaVersion, err := h.ipaVersionDAO.GetByIpaIDVersion(ctx, ipaID, args.Version)
	util.PanicIf(err)

	openURL, err := h.tencentCtl.GetSignatureURL(ctx, ipaVersion.TokenPath)
	util.PanicIf(err)
	res := map[string]interface{}{
		"open_url": openURL,
	}
	util.RenderJSON(w, res)
}
