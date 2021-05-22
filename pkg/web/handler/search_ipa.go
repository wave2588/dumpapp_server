package handler

import (
	"fmt"
	"net/http"

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
)

type SearchIpaHandler struct {
	ipaDAO          dao.IpaDAO
	ipaVersionDAO   dao.IpaVersionDAO
	searchRecordDAO dao.SearchRecordDAO

	emailWebCtl controller.EmailWebController
	appleCtl    controller2.AppleController
}

func NewSearchIpaHandler() *SearchIpaHandler {
	return &SearchIpaHandler{
		ipaDAO:          impl.DefaultIpaDAO,
		ipaVersionDAO:   impl.DefaultIpaVersionDAO,
		searchRecordDAO: impl.DefaultSearchRecordDAO,

		emailWebCtl: impl2.DefaultEmailWebController,
		appleCtl:    impl3.DefaultAppleController,
	}
}

type searchIpaArgs struct {
	Name  string `form:"name"`
	AppID int64  `form:"app_id"`
}

func (args *searchIpaArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *SearchIpaHandler) Search(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := searchIpaArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	loginID := mustGetLoginID(ctx)

	if args.Name == "" && args.AppID == 0 {
		panic(errors.UnproccessableError("请输入 app 名称或 app_id"))
	}

	/// 记录用户行为
	util.PanicIf(h.searchRecordDAO.Insert(ctx, &models.SearchRecord{
		MemberID: loginID,
		Keyword:  args.Name,
	}))

	ipaIDs := make([]int64, 0)
	if args.Name != "" {
		ids, err := h.ipaDAO.GetByLikeName(ctx, args.Name)
		util.PanicIf(err)
		ipaIDs = append(ipaIDs, ids...)
	}

	if args.AppID != 0 {
		appInfo, err := h.appleCtl.GetAppInfoByAppID(ctx, args.AppID)
		util.PanicIf(err)
		ids, err := h.ipaDAO.GetByLikeName(ctx, appInfo.Name)
		util.PanicIf(err)
		ipaIDs = append(ipaIDs, ids...)
	}

	data := render.NewIpaRender(ipaIDs, loginID, render.IpaDefaultRenderFields...).RenderSlice(ctx)
	if len(data) == 0 {
		util.PanicIf(h.emailWebCtl.SendEmailToMaster(ctx, args.Name, "dumpapp@126.com"))
	}
	util.RenderJSON(w, util.ListOutput{
		Paging: util.GenerateOffsetPaging(ctx, r, len(data), 0, len(data)),
		Data:   data,
	})
}
