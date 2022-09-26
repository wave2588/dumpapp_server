package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	"github.com/go-playground/validator/v10"
)

type AdminDispenseHandler struct {
	dispenseCountCtl controller.DispenseCountController
	accountDAO       dao.AccountDAO
}

func NewAdminDispenseHandler() *AdminDispenseHandler {
	return &AdminDispenseHandler{
		dispenseCountCtl: impl.DefaultDispenseCountController,
		accountDAO:       impl2.DefaultAccountDAO,
	}
}

type addDispenseCountArgs struct {
	Email  string `json:"email" validate:"required"`
	Number int64  `json:"number" validate:"required"`
}

func (p *addDispenseCountArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminDispenseHandler) AddCount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &addDispenseCountArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	accountMap, err := h.accountDAO.BatchGetByEmail(ctx, []string{args.Email})
	util.PanicIf(err)

	account, ok := accountMap[args.Email]
	if !ok {
		util.PanicIf(errors.ErrNotFoundMember)
		return
	}

	util.PanicIf(h.dispenseCountCtl.AddCount(ctx, account.ID, args.Number, enum.DispenseCountRecordTypeAdminPresented))

	util.RenderJSON(w, DefaultSuccessBody(ctx))
}

type deleteDispenseCountArgs struct {
	Email  string `json:"email" validate:"required"`
	Number int64  `json:"number" validate:"required"`
}

func (p *deleteDispenseCountArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminDispenseHandler) DeleteCount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &deleteDispenseCountArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	accountMap, err := h.accountDAO.BatchGetByEmail(ctx, []string{args.Email})
	util.PanicIf(err)

	account, ok := accountMap[args.Email]
	if !ok {
		util.PanicIf(errors.ErrNotFoundMember)
		return
	}
	util.PanicIf(h.dispenseCountCtl.DeductCount(ctx, account.ID, args.Number, 0, enum.DispenseCountRecordTypeAdminDeleted))

	util.RenderJSON(w, DefaultSuccessBody(ctx))
}
