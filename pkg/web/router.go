package web

import (
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/handler"
	"github.com/go-chi/chi"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	/// region config
	configHandler := handler.NewConfigHandler()
	r.With(middleware.OAuthGuest).Get("/config", configHandler.Get)
	// endregion

	/// region account
	accountHandler := handler.NewAccountHandler()
	r.With(middleware.OAuthGuest).Post("/email/captcha", accountHandler.SendEmailCaptcha)
	r.With(middleware.OAuthGuest).Post("/phone/captcha", accountHandler.SendPhoneCaptcha)
	r.With(middleware.OAuthGuest).Post("/register", accountHandler.Register)
	r.With(middleware.OAuthGuest).Post("/login", accountHandler.Login)
	r.With(middleware.OAuthGuest).Post("/logout", accountHandler.Logout)
	// endregion

	// member
	memberHandler := handler.NewMemberHandler()
	r.With(middleware.OAuthRegister).Get("/member/self", memberHandler.GetSelf)
	// endregion

	// ipa
	ipaHandler := handler.NewIpaHandler()
	r.With(middleware.OAuthRegister).Get("/ipa/{ipa_id}", ipaHandler.Get)
	r.With(middleware.OAuthGuest).Get("/ipa/ranking", ipaHandler.GetRanking)
	// endregion

	// email
	emailHandler := handler.NewEmailHandler()
	r.With(middleware.OAuthGuest).Post("/email", emailHandler.PostEmail)
	// endregion

	// region Cos
	tencentCosHandler := handler.NewTencentCosHandler()
	r.With(middleware.OAuthAdmin).Get("/cos", tencentCosHandler.Get)
	// endregion

	return r
}

var DefaultRouter = NewRouter()
