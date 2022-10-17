package open_api_handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	"github.com/go-playground/validator/v10"
)

type AuthWebsiteHandler struct {
	adminAuthWebsiteDAO dao.AdminAuthWebsiteDAO
}

func NewAuthWebsiteHandler() *AuthWebsiteHandler {
	return &AuthWebsiteHandler{
		adminAuthWebsiteDAO: impl.DefaultAdminAuthWebsiteDAO,
	}
}

type getAuthWebsiteArgs struct {
	Domain string `form:"domain"`
}

func (args *getAuthWebsiteArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

type getAuthWebsiteResponse struct {
	Success bool    `json:"success"`
	Message *string `json:"message,omitempty"`
}

func (h *AuthWebsiteHandler) GetAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := getAuthWebsiteArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	authWebsite, err := h.adminAuthWebsiteDAO.GetByDomainSafe(ctx, args.Domain)
	util.PanicIf(err)

	/// 未收录的站点
	if authWebsite == nil {
		util.RenderJSON(w, &getAuthWebsiteResponse{
			Success: false,
			Message: util.StringPtr("未授权的站点"),
		})
		return
	}

	/// 未打开权限的站点
	if !authWebsite.BizExt.IsOpen {
		util.RenderJSON(w, &getAuthWebsiteResponse{
			Success: false,
			Message: util.StringPtr("未授权的站点"),
		})
		return
	}

	util.RenderJSON(w, &getAuthWebsiteResponse{
		Success: true,
	})
}
