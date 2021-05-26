package web

import (
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/handler"
	"github.com/go-chi/chi"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	/// admin_v2
	adminIpaHandler := handler.NewAdminIpaHandler()
	r.With(middleware.OAuthGuest).Post("/ipa", adminIpaHandler.Post)

	/// region account
	accountHandler := handler.NewAccountHandler()
	r.With(middleware.OAuthGuest).Post("/captcha", accountHandler.SendCaptcha)
	r.With(middleware.OAuthGuest).Post("/register", accountHandler.Register)
	r.With(middleware.OAuthGuest).Post("/login", accountHandler.Login)
	r.With(middleware.OAuthGuest).Post("/logout", accountHandler.Logout)
	// endregion

	// member
	memberHandler := handler.NewMemberHandler()
	r.With(middleware.OAuthRegister).Get("/member/self", memberHandler.GetSelf)
	// endregion

	// member_vip
	memberVipHandler := handler.NewMemberVipHandler()
	r.With(middleware.OAuthRegister).Post("/member/vip", memberVipHandler.Post)
	r.With(middleware.OAuthGuest).Get("/member/vip", memberVipHandler.Get)
	// endregion

	// ipa
	ipaHandler := handler.NewIpaHandler()
	r.With(middleware.OAuthRegister).Get("/ipa", ipaHandler.List)
	r.With(middleware.OAuthRegister).Get("/ipa/{ipa_id}", ipaHandler.Get)
	// endregion

	// search ipa
	searchIpaHandler := handler.NewSearchIpaHandler()
	r.With(middleware.OAuthRegister).Post("/ipa/search", searchIpaHandler.Post)
	// endregion

	// email
	emailHandler := handler.NewEmailHandler()
	r.With(middleware.OAuthGuest).Post("/email", emailHandler.PostEmail)
	// endregion

	// region Cos
	tencentCosHandler := handler.NewTencentCosHandler()
	r.With().Get("/cos", tencentCosHandler.Get)
	// endregion

	// region alipay callback
	callbackPayHandler := handler.NewCallbackPayHandler()
	r.With().Post("/callback_alipay", callbackPayHandler.ALiPayCallback)
	// endregion

	return r
}

var DefaultRouter = NewRouter()
