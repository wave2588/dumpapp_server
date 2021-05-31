package handler

import (
	"fmt"
	"net/http"

	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/common/util"
	controller2 "dumpapp_server/pkg/controller"
	impl3 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/controller"
	impl2 "dumpapp_server/pkg/web/controller/impl"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	pkgErr "github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/volatiletech/null/v8"
)

type SearchIpaHandler struct {
	ipaDAO            dao.IpaDAO
	ipaVersionDAO     dao.IpaVersionDAO
	searchRecordV2DAO dao.SearchRecordV2DAO

	emailWebCtl controller.EmailWebController
	appleCtl    controller2.AppleController
	tencentCtl  controller2.TencentController
}

func NewSearchIpaHandler() *SearchIpaHandler {
	return &SearchIpaHandler{
		ipaDAO:            impl.DefaultIpaDAO,
		ipaVersionDAO:     impl.DefaultIpaVersionDAO,
		searchRecordV2DAO: impl.DefaultSearchRecordV2DAO,

		emailWebCtl: impl2.DefaultEmailWebController,
		appleCtl:    impl3.DefaultAppleController,
		tencentCtl:  impl3.DefaultTencentController,
	}
}

type postSearchArgs struct {
	IpaID   string `json:"ipa_id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Version string `json:"app_version"`
}

func (p *postSearchArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *SearchIpaHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := mustGetLoginID(ctx)
	account := GetAccountByLoginID(ctx, loginID)
	loginMember := render.NewMemberRender([]int64{loginID}, loginID, render.MemberDefaultRenderFields...).RenderMap(ctx)[loginID]
	if loginMember == nil {
		panic(errors.ErrNotFoundMember)
	}

	args := &postSearchArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	ipaID := cast.ToInt64(args.IpaID)
	if ipaID == 0 {
		panic(errors.HttpBadRequestError)
	}
	ipa, err := h.ipaDAO.Get(ctx, ipaID)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}

	util.PanicIf(h.searchRecordV2DAO.Insert(ctx, &models.SearchRecordV2{
		MemberID: loginID,
		IpaID:    ipaID,
		Name:     args.Name,
		Version:  null.StringFrom(args.Version),
	}))

	if ipa == nil {
		if loginMember.Vip.IsVip {
			util.PanicIf(h.emailWebCtl.SendVipEmailToMaster(ctx, args.Name, args.Version, account.Email))
		} else {
			util.PanicIf(h.emailWebCtl.SendEmailToMaster(ctx, args.Name, args.Version, account.Email))
		}
		panic(errors.ErrNotFoundIpa)
	}

	if args.Version != "" {
		/// 不是 vip 不能用此功能
		if !loginMember.Vip.IsVip {
			panic(errors.ErrUpgradeVip)
		}
		ipaVersion, err := h.ipaVersionDAO.GetByIpaIDVersion(ctx, ipaID, args.Version)
		if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
			panic(err)
		}
		if ipaVersion == nil {
			util.PanicIf(h.emailWebCtl.SendVipEmailToMaster(ctx, args.Name, args.Version, account.Email))
			panic(errors.ErrNotFoundIpaVersion)
		}
	}

	data := render.NewIpaRender([]int64{ipa.ID}, loginID, render.IpaDefaultRenderFields...).RenderSlice(ctx)
	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, len(data), 0, len(data)),
		Data:   data,
	})
}
