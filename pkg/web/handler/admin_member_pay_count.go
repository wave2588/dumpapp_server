package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/errors"
	"github.com/go-playground/validator/v10"
)

type AdminMemberPayCountHandler struct {
	accountDAO        dao.AccountDAO
	memberPayCountCtl controller.MemberPayCountController
}

func NewAdminMemberPayCountHandler() *AdminMemberPayCountHandler {
	return &AdminMemberPayCountHandler{
		accountDAO:        impl.DefaultAccountDAO,
		memberPayCountCtl: impl2.DefaultMemberPayCountController,
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

func (h *AdminMemberPayCountHandler) AddNumber(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &addDownloadNumber{}
	util.PanicIf(util.JSONArgs(r, args))

	accountMap, err := h.accountDAO.BatchGetByEmail(ctx, []string{args.Email})
	util.PanicIf(err)

	account, ok := accountMap[args.Email]
	if !ok {
		util.PanicIf(errors.ErrNotFoundMember)
		return
	}

	util.PanicIf(h.memberPayCountCtl.AddCount(ctx, account.ID, args.Number, enum.MemberPayCountSourceAdminPresented, datatype.MemberPayCountRecordBizExt{
		ObjectID:   0,
		ObjectType: datatype.MemberPayCountRecordBizExtObjectTypeNone,
	}))
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

func (h *AdminMemberPayCountHandler) DeleteNumber(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &deleteDownloadNumber{}
	util.PanicIf(util.JSONArgs(r, args))

	accountMap, err := h.accountDAO.BatchGetByEmail(ctx, []string{args.Email})
	util.PanicIf(err)

	account, ok := accountMap[args.Email]
	if !ok {
		util.PanicIf(errors.ErrNotFoundMember)
		return
	}

	util.PanicIf(h.memberPayCountCtl.CheckPayCount(ctx, account.ID, args.Number))

	util.PanicIf(h.memberPayCountCtl.DeductPayCount(ctx, account.ID, args.Number, enum.MemberPayCountStatusAdminDelete, enum.MemberPayCountUseAdminDelete, datatype.MemberPayCountRecordBizExt{
		ObjectID:   0,
		ObjectType: datatype.MemberPayCountRecordBizExtObjectTypeNone,
	}))
}
