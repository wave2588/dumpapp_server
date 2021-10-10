package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/controller"
	impl2 "dumpapp_server/pkg/web/controller/impl"
	"github.com/go-playground/validator/v10"
)

type FeedbackHandler struct {
	feedbackDAO dao.FeedbackDAO
	alertWebCtl controller.AlterWebController
}

func NewFeedbackHandler() *FeedbackHandler {
	return &FeedbackHandler{
		feedbackDAO: impl.DefaultFeedbackDAO,
		alertWebCtl: impl2.DefaultAlterWebController,
	}
}

type postFeedbackArgs struct {
	Content string `json:"content" validate:"required"`
}

func (p *postFeedbackArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *FeedbackHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)

	args := &postFeedbackArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	if len(args.Content) > 2000 {
		panic(errors.UnproccessableError("超过最大字数限制"))
	}

	util.PanicIf(h.feedbackDAO.Insert(ctx, &models.Feedback{
		MemberID: loginID,
		Content:  args.Content,
	}))

	h.alertWebCtl.SendFeedbackMsg(ctx, loginID, args.Content)

	util.RenderJSON(w, "ok")
}
