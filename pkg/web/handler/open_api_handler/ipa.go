package open_api_handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
)

type OpenIpaHandler struct {
	memberPayCountCtl controller.MemberPayCountController
}

func NewOpenIpaHandler() *OpenIpaHandler {
	return &OpenIpaHandler{
		memberPayCountCtl: impl.DefaultMemberPayCountController,
	}
}

type getIpaArgs struct {
	IpaID string `form:"ipa_id" validate:"required"`
}

func (p *getIpaArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *OpenIpaHandler) Get(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		loginID = mustGetLoginID(ctx, r)
	)

	args := getIpaArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	ipaID := cast.ToInt64(args.IpaID)
	ipaMap := render.NewIpaRender([]int64{ipaID}, loginID, []enum.IpaType{enum.IpaTypeNormal, enum.IpaTypeCrack}, render.IpaDefaultRenderFields...).RenderMap(ctx)
	ipa, ok := ipaMap[ipaID]
	if !ok {
		util.PanicIf(errors.ErrNotFoundApp)
	}
	util.RenderJSON(w, ipa)
}

type getIpaDownloadURLArgs struct {
	IpaID      string `form:"ipa_id" validate:"required"`
	IpaVersion string `form:"ipa_version" validate:"required"`
}

func (p *getIpaDownloadURLArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *OpenIpaHandler) GetIpaDownloadURL(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		loginID = mustGetLoginID(ctx, r)
	)
	util.PanicIf(h.memberPayCountCtl.CheckPayCount(ctx, loginID, constant.IpaPrice))
}
