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
)

type SearchIpaHandler struct {
	ipaDAO          dao.IpaDAO
	ipaVersionDAO   dao.IpaVersionDAO
	searchRecordDAO dao.SearchRecordDAO

	emailWebCtl controller.EmailWebController
	appleCtl    controller2.AppleController
	tencentCtl  controller2.TencentController
}

func NewSearchIpaHandler() *SearchIpaHandler {
	return &SearchIpaHandler{
		ipaDAO:          impl.DefaultIpaDAO,
		ipaVersionDAO:   impl.DefaultIpaVersionDAO,
		searchRecordDAO: impl.DefaultSearchRecordDAO,

		emailWebCtl: impl2.DefaultEmailWebController,
		appleCtl:    impl3.DefaultAppleController,
		tencentCtl:  impl3.DefaultTencentController,
	}
}

type postSearchArgs struct {
	AppID   int64  `json:"app_id" validate:"required"`
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
	_ = GetAccountByLoginID(ctx, loginID)
	loginMember := render.NewMemberRender([]int64{loginID}, loginID, render.MemberDefaultRenderFields...).RenderMap(ctx)[loginID]
	if loginMember == nil {
		panic(errors.ErrNotFoundMember)
	}

	args := &postSearchArgs{}
	util.PanicIf(util.JSONArgs(r, args))
	if args.Version != "" && !loginMember.Vip.IsVip {
		panic(errors.ErrUpgradeVip)
	}

	ipa, err := h.ipaDAO.Get(ctx, args.AppID)
	util.PanicIf(err)

	/// 记录用户行为
	util.PanicIf(h.searchRecordDAO.Insert(ctx, &models.SearchRecord{
		MemberID: loginID,
		Keyword:  ipa.Name,
	}))

	if err != nil && pkgErr.Cause(err) == errors2.ErrNotFound {
		util.PanicIf(h.emailWebCtl.SendEmailToMaster(ctx, ipa.Name, args.Version, loginMember.Email))
		util.RenderJSON(w, util.ListOutput{
			Paging: nil,
			Data:   []interface{}{},
		})
		return
	}
	util.PanicIf(err)

	if args.Version != "" {
		ipaVersion, err := h.ipaVersionDAO.GetByIpaIDVersion(ctx, ipa.ID, args.Version)
		if err != nil && pkgErr.Cause(err) == errors2.ErrNotFound {
			util.PanicIf(h.emailWebCtl.SendEmailToMaster(ctx, ipa.Name, args.Version, loginMember.Email))
			util.RenderJSON(w, util.ListOutput{
				Paging: nil,
				Data:   []interface{}{},
			})
			return
		}
		util.PanicIf(err)

		url, err := h.tencentCtl.GetSignatureURL(ctx, ipaVersion.TokenPath)
		util.PanicIf(err)
		data := []*render.Ipa{
			{
				ID:   ipa.ID,
				Name: ipa.Name,
				Versions: []*render.Version{
					{
						Version: args.Version,
						URL:     url,
					},
				},
			},
		}
		util.RenderJSON(w, util.ListOutput{
			Paging: util.GenerateOffsetPaging(ctx, r, len(data), 0, len(data)),
			Data:   data,
		})
		return
	}

	data := render.NewIpaRender([]int64{ipa.ID}, loginID, render.IpaDefaultRenderFields...).RenderSlice(ctx)
	if len(data) == 0 {
		util.PanicIf(h.emailWebCtl.SendEmailToMaster(ctx, ipa.Name, args.Version, loginMember.Email))
	}
	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, len(data), 0, len(data)),
		Data:   data,
	})
}
