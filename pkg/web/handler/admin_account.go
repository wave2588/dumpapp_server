package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	util2 "dumpapp_server/pkg/common/util"
	dao2 "dumpapp_server/pkg/dao"
	impl4 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/util"
	"dumpapp_server/pkg/web/controller"
	"dumpapp_server/pkg/web/controller/impl"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AdminAccountHandler struct {
	accountDAO            dao2.AccountDAO
	memberIDEncryptionDAO dao2.MemberIDEncryptionDAO
	alertWebCtl           controller.AlterWebController
}

func NewAdminAccountHandler() *AdminAccountHandler {
	return &AdminAccountHandler{
		accountDAO:            impl4.DefaultAccountDAO,
		memberIDEncryptionDAO: impl4.DefaultMemberIDEncryptionDAO,
		alertWebCtl:           impl.DefaultAlterWebController,
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

	_, err := h.getAccountByEmail(ctx, args.Email)
	if err != errors.ErrNotFoundMember {
		util2.PanicIf(err)
	}

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
		Code:     util.MustGenerateUUID(),
	}))

	util2.RenderJSON(w, "ok")
}

type putAccountArgs struct {
	Email      string           `json:"email"`
	NewEmail   *string          `json:"new_email"`
	Password   *string          `json:"password"`
	Phone      *string          `json:"phone"`
	Role       enum.AccountRole `json:"role"`
	RoleReason *string          `json:"role_reason"`
}

func (p *putAccountArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if p.Phone != nil && !constant.CheckPhoneValid(*p.Phone) {
		return errors.ErrPhoneRefusedRegister
	}
	if p.NewEmail != nil && !constant.CheckEmailValid(*p.NewEmail) {
		return errors.ErrEmailRefusedRegister
	}
	if p.Password != nil && len(*p.Password) < 8 {
		return errors.UnproccessableError("密码长度必须大于 8 位")
	}
	return nil
}

func (h *AdminAccountHandler) PutAccount(w http.ResponseWriter, r *http.Request) {
	var (
		ctx          = r.Context()
		loginAccount = mustGetLoginAccount(ctx)
	)

	args := &putAccountArgs{}
	util2.PanicIf(util2.JSONArgs(r, args))

	if args.NewEmail != nil {
		e := *args.NewEmail
		aMap, err := h.accountDAO.BatchGetByEmail(ctx, []string{e})
		util2.PanicIf(err)
		if _, ok := aMap[e]; ok {
			util2.PanicIf(errors.UnproccessableError("NewEmail 已存在"))
		}
	}
	if args.Phone != nil {
		p := *args.Phone
		aMap, err := h.accountDAO.BatchGetByPhones(ctx, []string{p})
		util2.PanicIf(err)
		if _, ok := aMap[p]; ok {
			util2.PanicIf(errors.UnproccessableError("Phone 已存在"))
		}
	}

	account, err := h.getAccountByEmail(ctx, args.Email)
	util2.PanicIf(err)

	/// 加个设置为代理的推送
	if args.RoleReason != nil && account.Role != args.Role {
		titleString := "<font color=\"warning\">设置代理商</font>\n>"
		countString := fmt.Sprintf("Reason：<font color=\"comment\">%s</font>\n", *args.RoleReason)
		receiveEmailString := fmt.Sprintf("用户邮箱：<font color=\"comment\">%s</font>\n", args.Email)
		adminEmailString := fmt.Sprintf("管理员邮箱：<font color=\"comment\">%s</font>\n", loginAccount.Email)
		timeStr := fmt.Sprintf("操作时间：<font color=\"comment\">%s</font>", time.Now().Format("2006-01-02 15:04:05"))
		h.alertWebCtl.SendCustomMsg(ctx, "32df4de7-524c-4d0c-94cd-c8d7e0709fb4", titleString+countString+receiveEmailString+adminEmailString+timeStr)
	}

	if args.NewEmail != nil {
		account.Email = *args.NewEmail
	}
	if args.Phone != nil {
		account.Phone = *args.Phone
	}
	if args.Password != nil {
		account.Password = *args.Password
	}
	if args.Role.IsAAccountRole() {
		account.Role = args.Role
	}
	util2.PanicIf(h.accountDAO.Update(ctx, account))

	util2.RenderJSON(w, "ok")
}

func (h *AdminAccountHandler) getAccountByEmail(ctx context.Context, email string) (*models.Account, error) {
	accountMap, err := h.accountDAO.BatchGetByEmail(ctx, []string{email})
	util2.PanicIf(err)
	account, ok := accountMap[email]
	if !ok {
		return nil, errors.ErrNotFoundMember
	}
	return account, nil
}

type accountListArgs struct {
	Role    enum.AccountRole `form:"role"`
	StartAt int64            `form:"start_at"`
	EndAt   int64            `form:"end_at"`
}

func (args *accountListArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminAccountHandler) AccountList(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		loginID = mustGetLoginID(ctx)
		offset  = GetIntArgument(r, "offset", 0)
		limit   = GetIntArgument(r, "limit", 10)
	)

	args := accountListArgs{}
	util2.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util2.PanicIf(args.Validate())

	filters := make([]qm.QueryMod, 0)
	if args.Role.IsAAccountRole() {
		filters = append(filters, models.AccountWhere.Role.EQ(args.Role))
	}
	if args.StartAt != 0 {
		filters = append(filters, models.AccountWhere.CreatedAt.GTE(cast.ToTime(args.StartAt)))
	}
	if args.EndAt != 0 {
		filters = append(filters, models.AccountWhere.CreatedAt.LTE(cast.ToTime(args.EndAt)))
	}

	ids, err := impl4.DefaultAccountDAO.ListIDs(ctx, offset, limit, filters, nil)
	util2.PanicIf(err)
	totalCount, err := impl4.DefaultAccountDAO.Count(ctx, filters)
	util2.PanicIf(err)

	members := render.NewMemberRender(ids, loginID, render.MemberAdminRenderFields...).RenderSlice(ctx)
	util2.RenderJSON(w, util2.ListOutput{
		Paging: util2.GenerateOffsetPaging(ctx, r, int(totalCount), offset, limit),
		Data:   members,
	})
}

type getAccountArgs struct {
	Account string `form:"account" validate:"required"`
}

func (args *getAccountArgs) Validate() error {
	err := validator.New().Struct(args)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AdminAccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		loginID = mustGetLoginID(ctx)
	)

	args := getAccountArgs{}
	util2.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util2.PanicIf(args.Validate())

	account := GetAccountByAccount(ctx, args.Account)
	if account == nil {
		util2.PanicIf(errors.ErrNotFoundMember)
	}

	accountID := account.ID
	memberMap := render.NewMemberRender([]int64{accountID}, loginID, render.MemberAdminRenderFields...).RenderMap(ctx)
	util2.RenderJSON(w, memberMap[accountID])
}
