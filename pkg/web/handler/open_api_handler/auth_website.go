package open_api_handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	"github.com/go-playground/validator/v10"
)

type AuthWebsiteHandler struct{}

func NewAuthWebsiteHandler() *AuthWebsiteHandler {
	return &AuthWebsiteHandler{}
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
	args := getAuthWebsiteArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	resp := &getAuthWebsiteResponse{}
	if util2.IsContainStrings(constant.AuthWebsites, args.Domain) {
		resp.Success = true
	} else {
		resp.Success = false
		resp.Message = util.StringPtr("未授权的站点")
	}
	util.RenderJSON(w, resp)
}
