package handler

import (
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type DailyFreeIpaHandler struct {
	dailyFreeDAO dao.DailyFreeRecordDAO
	configDAO    dao.AdminConfigDAO
	emailCtl     controller.EmailController
}

func NewDailyFreeIpaHandler() *DailyFreeIpaHandler {
	return &DailyFreeIpaHandler{
		dailyFreeDAO: impl.DefaultDailyFreeRecordDAO,
		configDAO:    impl.DefaultAdminConfigDAO,
		emailCtl:     impl2.DefaultEmailController,
	}
}

type postDailyFreeIpa struct {
	IpaID      int64  `json:"ipa_id,string" validate:"required"`
	IpaVersion string `json:"ipa_version" validate:"required"`
}

func (p *postDailyFreeIpa) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *DailyFreeIpaHandler) PostIpa(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)

	args := &postDailyFreeIpa{}
	util.PanicIf(util.JSONArgs(r, args))

	dailyFreeCount, err := h.configDAO.GetDailyFreeCount(ctx)
	util.PanicIf(err)

	now := time.Now()
	ids, err := h.dailyFreeDAO.ListIDs(ctx, 0, 100, []qm.QueryMod{
		models.DailyFreeRecordWhere.CreatedAt.GT(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())),
	}, nil)
	util.PanicIf(err)

	if len(ids) >= int(dailyFreeCount) {
		util.PanicIf(errors.HttpUnprocessableError)
		return
	}

	util.PanicIf(h.dailyFreeDAO.Insert(ctx, &models.DailyFreeRecord{
		MemberID:   loginID,
		IpaID:      args.IpaID,
		IpaVersion: args.IpaVersion,
	}))
}
