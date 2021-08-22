package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AdminDownloadNumberHandler struct {
	accountDAO              dao.AccountDAO
	memberDownloadNumberDAO dao.MemberDownloadNumberDAO
}

func NewAdminDownloadNumberHandler() *AdminDownloadNumberHandler {
	return &AdminDownloadNumberHandler{
		accountDAO:              impl.DefaultAccountDAO,
		memberDownloadNumberDAO: impl.DefaultMemberDownloadNumberDAO,
	}
}

type addDownloadNumber struct {
	Email  string `json:"email" validate:"required"`
	Number int64  `json:"number" validate:"required"`
}

func (p *addDownloadNumber) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminDownloadNumberHandler) AddNumber(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)
	if _, ok := constant.OpsAuthMemberIDMap[loginID]; !ok {
		panic(errors.ErrMemberAccessDenied)
	}

	args := &addDownloadNumber{}
	util.PanicIf(util.JSONArgs(r, args))

	account, err := h.accountDAO.GetByEmail(ctx, args.Email)
	util.PanicIf(err)

	for i := 0; i < cast.ToInt(args.Number); i++ {
		util.PanicIf(h.memberDownloadNumberDAO.Insert(ctx, &models.MemberDownloadNumber{
			MemberID: account.ID,
			Status:   enum.MemberDownloadNumberStatusNormal,
		}))
	}
}

type deleteDownloadNumber struct {
	Email  string `json:"email" validate:"required"`
	Number int64  `json:"number" validate:"required"`
}

func (p *deleteDownloadNumber) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminDownloadNumberHandler) DeleteNumber(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loginID := middleware.MustGetMemberID(ctx)
	if _, ok := constant.OpsAuthMemberIDMap[loginID]; !ok {
		panic(errors.ErrMemberAccessDenied)
	}

	args := &deleteDownloadNumber{}
	util.PanicIf(util.JSONArgs(r, args))

	account, err := h.accountDAO.GetByEmail(ctx, args.Email)
	util.PanicIf(err)

	filter := []qm.QueryMod{
		models.MemberDownloadNumberWhere.MemberID.EQ(account.ID),
		models.MemberDownloadNumberWhere.Status.EQ(enum.MemberDownloadNumberStatusNormal),
	}
	dnIDs, err := h.memberDownloadNumberDAO.ListIDs(ctx, 0, 10000, filter, nil)
	util.PanicIf(err)

	for i := 0; i < cast.ToInt(args.Number); i++ {
		if i >= len(dnIDs) {
			break
		}
		id := dnIDs[i]
		util.PanicIf(h.memberDownloadNumberDAO.Delete(ctx, id))
	}
}
