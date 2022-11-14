package handler

import (
	"context"
	errors3 "errors"
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/clients"
	"dumpapp_server/pkg/common/constant"
	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/common/util"
	controller2 "dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	dao2 "dumpapp_server/pkg/dao"
	impl4 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/middleware"
	util2 "dumpapp_server/pkg/middleware/util"
	util3 "dumpapp_server/pkg/util"
	"dumpapp_server/pkg/web/controller"
	"dumpapp_server/pkg/web/controller/impl"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	pkgErr "github.com/pkg/errors"
)

type AccountHandler struct {
	accountDAO            dao2.AccountDAO
	captchaDAO            dao2.CaptchaDAO
	memberInviteCodeDAO   dao2.MemberInviteCodeDAO
	memberInviteDAO       dao2.MemberInviteDAO
	memberIDEncryptionDAO dao2.MemberIDEncryptionDAO
	memberPayCountDAO     dao2.MemberPayCountDAO

	emailCtl    controller2.EmailController
	tencentCtl  controller2.TencentController
	alertWebCtl controller.AlterWebController
}

func NewAccountHandler() *AccountHandler {
	return &AccountHandler{
		accountDAO:            impl4.DefaultAccountDAO,
		captchaDAO:            impl4.DefaultCaptchaDAO,
		memberInviteCodeDAO:   impl4.DefaultMemberInviteCodeDAO,
		memberInviteDAO:       impl4.DefaultMemberInviteDAO,
		memberIDEncryptionDAO: impl4.DefaultMemberIDEncryptionDAO,
		memberPayCountDAO:     impl4.DefaultMemberPayCountDAO,

		emailCtl:    impl2.DefaultEmailController,
		tencentCtl:  impl2.DefaultTencentController,
		alertWebCtl: impl.DefaultAlterWebController,
	}
}

type sendEmailCaptchaQueryArgs struct {
	Email string `json:"email" validate:"required"`
}

func (p *sendEmailCaptchaQueryArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AccountHandler) SendEmailCaptcha(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &sendEmailCaptchaQueryArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	/// 检测邮箱
	if !constant.CheckEmailValid(args.Email) {
		panic(errors.ErrEmailRefusedRegister)
	}

	accountMap, err := h.accountDAO.BatchGetByEmail(ctx, []string{args.Email})
	util.PanicIf(err)
	account := accountMap[args.Email]
	if account != nil {
		panic(errors.ErrAccountRegisteredByEmail)
		return
	}

	captcha, err := h.captchaDAO.GetEmailCaptcha(ctx, args.Email)
	util.PanicIf(err)

	if captcha != "" {
		panic(errors.UnproccessableError("重复发送验证码"))
	}

	/// 发送验证码
	util.PanicIf(h.sendCaptcha(ctx, args.Email))
	util.RenderJSON(w, "ok")
}

func (h *AccountHandler) sendCaptcha(ctx context.Context, email string) error {
	captcha := util3.MustGenerateCaptcha(ctx)
	err := h.emailCtl.SendRegisterEmail(ctx, "验证码来了~", fmt.Sprintf("欢迎注册 iOS 脱壳平台, 此次注册验证码为: %s, 有效期为 5 分钟.", captcha), email)
	if err != nil {
		return err
	}
	return h.captchaDAO.SetEmailCaptcha(ctx, email, captcha)
}

type sendPhoneCaptchaQueryArgs struct {
	Phone string `json:"phone" validate:"required"`
}

func (p *sendPhoneCaptchaQueryArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AccountHandler) SendPhoneCaptcha(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &sendPhoneCaptchaQueryArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	/// 检测手机号
	if !constant.CheckPhoneValid(args.Phone) {
		panic(errors.ErrPhoneRefusedRegister)
	}

	accountMap, err := h.accountDAO.BatchGetByPhones(ctx, []string{args.Phone})
	util.PanicIf(err)
	account := accountMap[args.Phone]
	if account != nil {
		panic(errors.ErrAccountRegisteredByPhone)
		return
	}

	captcha, err := h.captchaDAO.GetPhoneCaptcha(ctx, args.Phone)
	util.PanicIf(err)

	if captcha != "" {
		panic(errors.ErrCaptchaRepeated)
	}

	/// 发送验证码
	newCaptcha := util3.MustGenerateCaptcha(ctx)
	util.PanicIf(h.captchaDAO.SetPhoneCaptcha(ctx, args.Phone, newCaptcha))
	util.PanicIf(h.tencentCtl.SendPhoneRegisterCaptcha(ctx, newCaptcha, args.Phone))

	util.RenderJSON(w, "ok")
}

type registerQueryArgs struct {
	Email        string `json:"email" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	PhoneCaptcha string `json:"phone_captcha" validate:"required"`
	Password     string `json:"password" validate:"required"`
	InviteCode   string `json:"invite_code"`
}

func (p *registerQueryArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if len(p.Password) < 8 {
		panic(errors.UnproccessableError("密码长度必须大于 8 位"))
	}
	return nil
}

func (h *AccountHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &registerQueryArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	/// 检测邮箱
	if !constant.CheckEmailValid(args.Email) {
		panic(errors.ErrEmailRefusedRegister)
	}
	/// 检测手机号
	if !constant.CheckPhoneValid(args.Phone) {
		panic(errors.ErrPhoneRefusedRegister)
	}

	account, err := h.accountDAO.GetByEmail(ctx, args.Email)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}
	if account != nil {
		panic(errors.ErrAccountRegisteredByEmail)
	}

	/// 验证手机号是否可用
	phoneCaptcha, err := h.captchaDAO.GetPhoneCaptcha(ctx, args.Phone)
	util.PanicIf(err)
	if args.PhoneCaptcha != phoneCaptcha {
		panic(errors.ErrCaptchaIncorrectByPhone)
	}
	accountMap, err := h.accountDAO.BatchGetByPhones(ctx, []string{args.Phone})
	util.PanicIf(err)
	account = accountMap[args.Phone]
	if account != nil {
		panic(errors.ErrAccountRegisteredByPhone)
	}

	accountID := util3.MustGenerateID(ctx)

	/// 生成邀请码, 并检查邀请码是否存在
	inviteCodeString := util3.MustGenerateCode(ctx, 10)
	inviteCode, err := h.memberInviteCodeDAO.GetByCode(ctx, inviteCodeString)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}
	if inviteCode != nil {
		util.PanicIf(errors3.New(fmt.Sprintf("邀请码已经存在. email: %s  phone: %s  invite_code_string: %s", args.Email, args.Phone, inviteCodeString)))
	}

	/// 事物
	txn := clients.GetMySQLTransaction(ctx, clients.MySQLConnectionsPool, true)
	defer clients.MustClearMySQLTransaction(ctx, txn)
	ctx = context.WithValue(ctx, constant.TransactionKeyTxn, txn)

	util.PanicIf(h.captchaDAO.RemoveEmailCaptcha(ctx, args.Email))
	util.PanicIf(h.captchaDAO.RemovePhoneCaptcha(ctx, args.Phone))

	/// start 判断是否有邀请码
	if args.InviteCode != "" {
		inviteCode, err := h.memberInviteCodeDAO.GetByCode(ctx, args.InviteCode)
		if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
			util.PanicIf(err)
		}
		if inviteCode == nil {
			panic(errors.ErrMemberInviteCodeInvalid)
		}

		/// 记录邀请关系
		util.PanicIf(h.memberInviteDAO.Insert(ctx, &models.MemberInvite{
			InviterID: inviteCode.MemberID,
			InviteeID: accountID,
		}))

		/// 邀请一个人只送 2 个币
		//for i := 0; i < 2; i++ {
		//util.PanicIf(h.memberPayCountDAO.Insert(ctx, &models.MemberPayCount{
		//	MemberID: inviteCode.MemberID,
		//	Status:   enum.MemberPayCountStatusNormal,
		//	Source:   enum.MemberPayCountSourceInvitedPresented,
		//}))
		//}
	}
	/// end

	util.PanicIf(h.accountDAO.Insert(ctx, &models.Account{
		ID:       accountID,
		Email:    args.Email,
		Password: args.Password,
		Phone:    args.Phone,
	}))

	util.PanicIf(h.memberIDEncryptionDAO.Insert(ctx, &models.MemberIDEncryption{
		MemberID: accountID,
		Code:     util3.MustGenerateUUID(),
	}))

	/// 写入邀请码
	util.PanicIf(h.memberInviteCodeDAO.Insert(ctx, &models.MemberInviteCode{
		MemberID: accountID,
		Code:     inviteCodeString,
	}))

	clients.MustCommit(ctx, txn)
	ctx = util.ResetCtxKey(ctx, constant.TransactionKeyTxn)

	members := render.NewMemberRender([]int64{accountID}, accountID, render.MemberDefaultRenderFields...).RenderSlice(ctx)

	/// 获取 ticket
	ticket, err := util2.GenerateRegisterTicket(accountID)
	util.PanicIf(err)
	middleware.SetTicketCookie(w, r, ticket)

	h.alertWebCtl.SendAccountMsg(ctx)

	util.RenderJSON(w, members[0])
}

type loginQueryArgs struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (p *loginQueryArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if len(p.Password) < 8 {
		panic(errors.UnproccessableError("密码长度必须大于 8 位"))
	}
	return nil
}

func (h *AccountHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &loginQueryArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	account := GetAccountByAccount(ctx, args.Account)
	if account == nil {
		panic(errors.ErrNotFoundMember)
	}

	if account.Password != args.Password {
		panic(errors.UnproccessableError("密码错误"))
	}
	members := render.NewMemberRender([]int64{account.ID}, account.ID, render.MemberDefaultRenderFields...).RenderSlice(ctx)

	/// 获取 ticket
	ticket, err := util2.GenerateRegisterTicket(account.ID)
	util.PanicIf(err)
	middleware.SetTicketCookie(w, r, ticket)

	util.RenderJSON(w, members[0])
}

func (h *AccountHandler) Logout(w http.ResponseWriter, r *http.Request) {
	util.ClearCookie(w, "session")
	util.RenderJSON(w, "ok")
}

type sendResetPasswordCaptchaQueryArgs struct {
	Email string `json:"email" validate:"required"`
}

func (p *sendResetPasswordCaptchaQueryArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AccountHandler) SendResetPasswordCaptcha(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &sendResetPasswordCaptchaQueryArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	/// 检测邮箱
	if !constant.CheckEmailValid(args.Email) {
		panic(errors.ErrEmailRefusedRegister)
	}

	accountMap, err := h.accountDAO.BatchGetByEmail(ctx, []string{args.Email})
	util.PanicIf(err)
	account := accountMap[args.Email]
	if account == nil {
		panic(errors.ErrNotFoundMember)
		return
	}

	newCaptcha := util3.MustGenerateCaptcha(ctx)
	util.PanicIf(h.emailCtl.SendRegisterEmail(ctx, "验证码来了~", fmt.Sprintf("DumpApp - 您正在修改账号密码，验证码为：%s，15 分钟有效，为保障帐户安全，请勿向任何人提供此验证码。\n", newCaptcha), args.Email))
	util.PanicIf(h.captchaDAO.SetResetPassowordCaptcha(ctx, args.Email, newCaptcha))

	util.RenderJSON(w, DefaultSuccessBody(ctx))
}

type resetPasswordQueryArgs struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Captcha  string `json:"captcha" validate:"required"`
}

func (p *resetPasswordQueryArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	if len(p.Password) < 8 {
		panic(errors.UnproccessableError("密码长度必须大于 8 位"))
	}
	return nil
}

func (h *AccountHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &resetPasswordQueryArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	/// 检测邮箱
	if !constant.CheckEmailValid(args.Email) {
		panic(errors.ErrEmailRefusedRegister)
	}

	accountMap, err := h.accountDAO.BatchGetByEmail(ctx, []string{args.Email})
	util.PanicIf(err)
	account := accountMap[args.Email]
	if account == nil {
		panic(errors.ErrNotFoundMember)
		return
	}

	captcha, err := h.captchaDAO.GetResetPassowordCaptcha(ctx, args.Email)
	util.PanicIf(err)

	if args.Captcha != captcha {
		panic(errors.UnproccessableError("验证码错误"))
	}

	_ = h.captchaDAO.RemoveResetPassowordCaptcha(ctx, args.Email)

	account.Password = args.Password
	util.PanicIf(h.accountDAO.Update(ctx, account))

	util.RenderJSON(w, DefaultSuccessBody(ctx))
}

type resetEmailQueryArgs struct {
	Email string `json:"email" validate:"required"`
}

func (p *resetEmailQueryArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AccountHandler) ResetEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)

	args := &resetEmailQueryArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	/// 检测邮箱
	if !constant.CheckEmailValid(args.Email) {
		panic(errors.ErrEmailRefusedRegister)
	}

	/// 判断邮箱是否存在
	emailAccountMap, err := h.accountDAO.BatchGetByEmail(ctx, []string{args.Email})
	util.PanicIf(err)
	_, ok := emailAccountMap[args.Email]
	if ok {
		util.PanicIf(errors.ErrAccountRegisteredByEmail)
	}

	/// 获取原本账号信息，并且修改密码
	accountMap, err := h.accountDAO.BatchGet(ctx, []int64{loginID})
	util.PanicIf(err)
	account, ok := accountMap[loginID]
	if !ok {
		util.PanicIf(errors.ErrNotFoundMember)
	}

	account.Email = args.Email
	util.PanicIf(h.accountDAO.Update(ctx, account))

	util.RenderJSON(w, DefaultSuccessBody(ctx))
}

func (h *AccountHandler) ResetToken(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		loginID = mustGetLoginID(ctx)
	)

	mtMap, err := h.memberIDEncryptionDAO.BatchGetByMemberID(ctx, []int64{loginID})
	util.PanicIf(err)

	if mt, ok := mtMap[loginID]; ok {
		mt.Code = util3.MustGenerateUUID()
		util.PanicIf(h.memberIDEncryptionDAO.Update(ctx, mt))
	} else {
		util.PanicIf(h.memberIDEncryptionDAO.Insert(ctx, &models.MemberIDEncryption{
			MemberID: loginID,
			Code:     util3.MustGenerateUUID(),
		}))
	}

	util.RenderJSON(w, DefaultSuccessBody(ctx))
}
