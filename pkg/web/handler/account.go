package handler

import (
	"context"
	"dumpapp_server/pkg/common/enum"
	"fmt"
	"net/http"

	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/common/util"
	controller2 "dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	dao2 "dumpapp_server/pkg/dao"
	impl4 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	rpc "dumpapp_server/pkg/ice"
	"dumpapp_server/pkg/ice/impl"
	"dumpapp_server/pkg/middleware"
	util2 "dumpapp_server/pkg/middleware/util"
	"dumpapp_server/pkg/web/render"
	"github.com/go-playground/validator/v10"
	pkgErr "github.com/pkg/errors"
)

type AccountHandler struct {
	iceRPC rpc.IceRPC

	accountDAO              dao2.AccountDAO
	captchaDAO              dao2.CaptchaDAO
	memberDownloadNumberDAO dao2.MemberDownloadNumberDAO

	emailCtl controller2.EmailController
}

func NewAccountHandler() *AccountHandler {
	return &AccountHandler{
		iceRPC: impl.DefaultIceRPC,

		accountDAO:              impl4.DefaultAccountDAO,
		captchaDAO:              impl4.DefaultCaptchaDAO,
		memberDownloadNumberDAO: impl4.DefaultMemberDownloadNumberDAO,

		emailCtl: impl2.DefaultEmailController,
	}
}

type sendCaptchaQueryArgs struct {
	Email string `json:"email"`
}

func (p *sendCaptchaQueryArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *AccountHandler) SendCaptcha(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &sendCaptchaQueryArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	captcha, err := h.captchaDAO.GetCaptcha(ctx, args.Email)
	util.PanicIf(err)

	if captcha != "" {
		panic(errors.UnproccessableError("重复发送验证码"))
	}

	/// 发送验证码
	util.PanicIf(h.sendCaptcha(ctx, args.Email))
	util.RenderJSON(w, "ok")
}

func (h *AccountHandler) sendCaptcha(ctx context.Context, email string) error {
	captcha := h.iceRPC.MustGenerateCaptcha(ctx)
	err := h.emailCtl.SendEmail(ctx, "验证码来了~", fmt.Sprintf("欢迎注册 iOS 脱壳平台, 此次注册验证码为: %s, 有效期为 5 分钟.", captcha), email, []string{})
	if err != nil {
		return err
	}
	return h.captchaDAO.SetCaptcha(ctx, email, captcha)
}

type registerQueryArgs struct {
	Email    string `json:"email"`
	Captcha  string `json:"captcha"`
	Password string `json:"password"`
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

	captcha, err := h.captchaDAO.GetCaptcha(ctx, args.Email)
	util.PanicIf(err)

	if args.Captcha != captcha {
		panic(errors.UnproccessableError("验证码错误"))
	}

	account, err := h.accountDAO.GetByEmail(ctx, args.Email)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}
	if account != nil {
		panic(errors.UnproccessableError("该邮箱已注册"))
	}

	accountID := h.iceRPC.MustGenerateID(ctx)
	util.PanicIf(h.accountDAO.Insert(ctx, &models.Account{
		ID:       accountID,
		Email:    args.Email,
		Password: args.Password,
	}))

	util.PanicIf(h.captchaDAO.RemoveCaptcha(ctx, args.Email))

	/// 新用户送一次下载次数
	util.PanicIf(h.memberDownloadNumberDAO.Insert(ctx, &models.MemberDownloadNumber{
		MemberID: accountID,
		Status:   enum.MemberDownloadNumberStatusNormal,
	}))

	members := render.NewMemberRender([]int64{accountID}, 0, render.MemberDefaultRenderFields...).RenderSlice(ctx)

	/// 获取 ticket
	ticket, err := util2.GenerateRegisterTicket(accountID)
	util.PanicIf(err)
	middleware.SetTicketCookie(w, r, ticket)

	util.RenderJSON(w, members[0])
}

type loginQueryArgs struct {
	Email    string `json:"email"`
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

	account := GetAccountByEmail(ctx, args.Email)

	if account.Password != args.Password {
		panic(errors.UnproccessableError("密码错误"))
	}
	members := render.NewMemberRender([]int64{account.ID}, 0, render.MemberDefaultRenderFields...).RenderSlice(ctx)

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
