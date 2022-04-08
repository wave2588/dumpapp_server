package handler

import (
	"fmt"
	"net/http"

	util2 "dumpapp_server/pkg/common/util"
	dao2 "dumpapp_server/pkg/dao"
	impl4 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/util"
	"github.com/go-playground/validator/v10"
)

type AdminAccountHandler struct {
	accountDAO            dao2.AccountDAO
	memberIDEncryptionDAO dao2.MemberIDEncryptionDAO
}

func NewAdminAccountHandler() *AdminAccountHandler {
	return &AdminAccountHandler{
		accountDAO:            impl4.DefaultAccountDAO,
		memberIDEncryptionDAO: impl4.DefaultMemberIDEncryptionDAO,
	}
}

type addAccountArgs struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (p *addAccountArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminAccountHandler) AddAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &addAccountArgs{}
	util2.PanicIf(util2.JSONArgs(r, args))

	id := util.MustGenerateID(ctx)
	util2.PanicIf(h.accountDAO.Insert(ctx, &models.Account{
		ID:       id,
		Email:    args.Email,
		Phone:    args.Phone,
		Password: args.Password,
		Status:   0,
	}))

	util2.PanicIf(h.memberIDEncryptionDAO.Insert(ctx, &models.MemberIDEncryption{
		MemberID: id,
		Code:     util.MustGenerateCode(ctx, 10),
	}))

	util2.RenderJSON(w, "ok")
}

type putAccountArgs struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p *putAccountArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminAccountHandler) PutAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &putAccountArgs{}
	util2.PanicIf(util2.JSONArgs(r, args))

	account, err := h.accountDAO.GetByEmail(ctx, args.Email)
	util2.PanicIf(err)

	account.Password = args.Password
	util2.PanicIf(h.accountDAO.Update(ctx, account))

	util2.RenderJSON(w, "ok")
}
