package web

import (
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/handler"
	"github.com/go-chi/chi"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	/// check health
	checkHealthHandler := handler.NewCheckHealthHandler()
	r.With(middleware.OAuthGuest).Post("/check_health", checkHealthHandler.Get)
	r.With(middleware.OAuthGuest).Get("/check_health", checkHealthHandler.Get)
	r.With(middleware.OAuthGuest).Put("/check_health", checkHealthHandler.Get)
	r.With(middleware.OAuthGuest).Delete("/check_health", checkHealthHandler.Get)
	/// endregion

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
	r.With(middleware.OAuthRegister).Get("/ipa/{ipa_id}/latest", ipaHandler.GetLatestVersion)
	r.With(middleware.OAuthGuest).Get("/ipa/ranking", ipaHandler.GetRanking)
	// endregion

	// dump order
	dumpOrderHandler := handler.NewDumpOrderHandler()
	r.With(middleware.OAuthRegister).Get("/ipa/dump_order", dumpOrderHandler.Post)
	// endregion

	// device
	deviceHandler := handler.NewDeviceHandler()
	/// 获取"获取"描述文件接口
	r.With(middleware.OAuthRegister).Get("/device/config/qr_code", deviceHandler.GetMobileConfigQRCode)
	/// 获取描述文件
	r.With(middleware.OAuthGuest).Get("/device/config/file", deviceHandler.GetMobileConfigFile)
	/// 绑定设备
	r.With(middleware.OAuthGuest).Post("/device/bind/{code}", deviceHandler.Bind)
	// endregion

	// certificate
	certificateHandler := handler.NewCertificateHandler()
	r.With(middleware.OAuthRegister).Post("/certificate", certificateHandler.Post) /// 生成证书
	r.With(middleware.OAuthRegister).Get("/certificate/p12", certificateHandler.DownloadP12File)
	r.With(middleware.OAuthRegister).Get("/certificate/mobileprovision", certificateHandler.DownloadMobileprovisionFile)
	// endregion

	// email
	emailHandler := handler.NewEmailHandler()
	r.With(middleware.OAuthGuest).Post("/email", emailHandler.PostEmail)
	// endregion

	// region feedback
	feedbackHandler := handler.NewFeedbackHandler()
	r.With(middleware.OAuthRegister).Post("/feedback", feedbackHandler.Post)
	// endregion

	// region Cos
	tencentCosHandler := handler.NewTencentCosHandler()
	r.With(middleware.OAuthAdmin).Get("/cos", tencentCosHandler.Get)
	// endregion

	return r
}

var DefaultRouter = NewRouter()
