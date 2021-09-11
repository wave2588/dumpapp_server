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
	r.With(middleware.OAuthRegister).Post("/admin/ipa", adminIpaHandler.Post)
	r.With(middleware.OAuthRegister).Delete("/admin/ipa", adminIpaHandler.DeleteIpa)
	r.With(middleware.OAuthRegister).Delete("/admin/batch_ipa", adminIpaHandler.BatchDeleteIpa)
	// endregion

	/// admin_record
	adminSearchRecordHandler := handler.NewAdminSearchRecordHandler()
	r.With(middleware.OAuthAdmin).Get("/admin/search/record", adminSearchRecordHandler.GetMemberSearchRecord)

	/// admin download number
	adminDownloadNumberHandler := handler.NewAdminDownloadNumberHandler()
	r.With(middleware.OAuthRegister).Post("/admin/member/download_number", adminDownloadNumberHandler.AddNumber)
	r.With(middleware.OAuthRegister).Delete("/admin/member/download_number", adminDownloadNumberHandler.DeleteNumber)

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
	// endregion

	// email
	emailHandler := handler.NewEmailHandler()
	r.With(middleware.OAuthGuest).Post("/email", emailHandler.PostEmail)
	// endregion

	// region Cos
	tencentCosHandler := handler.NewTencentCosHandler()
	r.With().Get("/cos", tencentCosHandler.Get)
	// endregion

	return r
}

var DefaultRouter = NewRouter()
