package handler

import (
	"context"
	"dumpapp_server/pkg/common/datatype"
	util2 "dumpapp_server/pkg/util"
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/clients"
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

type AdminMemberPayCountHandler struct {
	accountDAO                dao.AccountDAO
	memberPayCountDAO         dao.MemberPayCountDAO
	memberPayExpenseRecordDAO dao.MemberPayExpenseRecordDAO
}

func NewAdminMemberPayCountHandler() *AdminMemberPayCountHandler {
	return &AdminMemberPayCountHandler{
		accountDAO:                impl.DefaultAccountDAO,
		memberPayCountDAO:         impl.DefaultMemberPayCountDAO,
		memberPayExpenseRecordDAO: impl.DefaultMemberPayExpenseRecordDAO,
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
	loginID := middleware.MustGetMemberID(ctx)
	if _, ok := constant.OpsAuthMemberIDMap[loginID]; !ok {
		panic(errors.ErrMemberAccessDenied)
	}

	args := &addDownloadNumber{}
	util.PanicIf(util.JSONArgs(r, args))

	account, err := h.accountDAO.GetByEmail(ctx, args.Email)
	util.PanicIf(err)

	for i := 0; i < cast.ToInt(args.Number); i++ {
		util.PanicIf(h.memberPayCountDAO.Insert(ctx, &models.MemberPayCount{
			MemberID: account.ID,
			Status:   enum.MemberPayCountStatusNormal,
			Source:   enum.MemberPayCountSourceAdminPresented,
		}))
	}

	/// 写入充值记录
	util.PanicIf(h.memberPayExpenseRecordDAO.Insert(ctx, &models.MemberPayExpenseRecord{
		MemberID: account.ID,
		Status:   enum.MemberPayExpenseRecordStatusAdd,
		Count:    args.Number,
		BizExt: datatype.MemberPayExpenseRecordBizExt{
			CountSource: enum.MemberPayCountSourceAdminPresented,
		},
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
	loginID := middleware.MustGetMemberID(ctx)
	if _, ok := constant.OpsAuthMemberIDMap[loginID]; !ok {
		panic(errors.ErrMemberAccessDenied)
	}

	args := &deleteDownloadNumber{}
	util.PanicIf(util.JSONArgs(r, args))

	account, err := h.accountDAO.GetByEmail(ctx, args.Email)
	util.PanicIf(err)

	filter := []qm.QueryMod{
		models.MemberPayCountWhere.MemberID.EQ(account.ID),
		models.MemberPayCountWhere.Status.EQ(enum.MemberPayCountStatusNormal),
	}

	ids, err := h.memberPayCountDAO.ListIDs(ctx, 0, int(args.Number), filter, nil)
	util.PanicIf(err)
	pcMap, err := h.memberPayCountDAO.BatchGet(ctx, ids)
	util.PanicIf(err)

	if int(args.Number) != len(ids) {
		util.PanicIf(errors.UnproccessableError("没有足够的次数可以扣除"))
		return
	}

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)
	for _, count := range pcMap {
		count.Status = enum.MemberPayCountStatusAdminDelete
		util.PanicIf(h.memberPayCountDAO.Update(ctx, count))
	}

	util.PanicIf(h.memberPayExpenseRecordDAO.Insert(ctx, &models.MemberPayExpenseRecord{
		MemberID: account.ID,
		Status:   enum.MemberPayExpenseRecordStatusReduce,
		Count:    args.Number,
		BizExt: datatype.MemberPayExpenseRecordBizExt{
			AdminMemberID: util2.Int64Ptr(loginID),
		},
	}))

	clients.MustCommit(ctx, txn)
	ctx = util.ResetCtxKey(ctx, constant.TransactionKeyTxn)
}
