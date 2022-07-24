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
	r.With(middleware.OAuthGuest).Post("/reset/password/captcha", accountHandler.SendResetPasswordCaptcha)
	r.With(middleware.OAuthGuest).Post("/reset/password", accountHandler.ResetPassword)
	r.With(middleware.OAuthRegister).Post("/reset/email", accountHandler.ResetEmail)
	// endregion

	// member
	memberHandler := handler.NewMemberHandler()
	r.With(middleware.OAuthRegister).Get("/member/self", memberHandler.GetSelf)
	r.With(middleware.OAuthRegister).Get("/member/self/devices", memberHandler.GetSelfDevice)
	r.With(middleware.OAuthRegister).Get("/member/self/devices_v2", memberHandler.GetSelfDeviceV2)
	r.With(middleware.OAuthRegister).Get("/member/self/certificates", memberHandler.GetSelfCertificate)
	r.With(middleware.OAuthRegister).Get("/member/self/coin/records", memberHandler.GetSelfCoinRecords)
	// endregion

	// member download record
	memberDownloadRecord := handler.NewMemberDownloadRecordHandler()
	r.With(middleware.OAuthRegister).Get("/member/self/download_record", memberDownloadRecord.GetSelfDownloadRecord)
	// endregion

	// member rebate record
	memberRebateRecord := handler.NewMemberRebateRecordHandler()
	r.With(middleware.OAuthRegister).Get("/member/self/rebate_record", memberRebateRecord.GetRebateRecords)
	// endregion

	// ipa
	ipaHandler := handler.NewIpaHandler()
	r.With(middleware.OAuthRegister).Get("/ipa/{ipa_id}", ipaHandler.Get)
	r.With(middleware.OAuthRegister).Get("/ipa/{ipa_id}/latest", ipaHandler.GetLatestVersion)
	r.With(middleware.OAuthRegister).Get("/ipa/{country}/{ipa_id}", ipaHandler.GetAllVersion)
	r.With(middleware.OAuthGuest).Get("/ipa/ranking", ipaHandler.GetRanking)
	// endregion

	// ipa list
	ipaListHandler := handler.NewIpaListHandler()
	r.With(middleware.OAuthGuest).Get("/ipa/{ipa_type}/list", ipaListHandler.GetByIpaType)
	// endregion

	// ipa sign
	ipaSignHandler := handler.NewIpaSignHandler()
	r.With(middleware.OAuthRegister).Post("/ipa/sign", ipaSignHandler.PostSign)
	r.With(middleware.OAuthRegister).Get("/ipa/sign", ipaSignHandler.GetMemberSignList)
	r.With(middleware.OAuthRegister).Get("/ipa/sign/{ipa_sign_id}/url", ipaSignHandler.GetIpaSignURL)
	// endregion

	// dump order
	dumpOrderHandler := handler.NewDumpOrderHandler()
	r.With(middleware.OAuthRegister).Post("/ipa/dump_order", dumpOrderHandler.Post)
	r.With(middleware.OAuthGuest).Get("/ipa/dump_order/list", dumpOrderHandler.GetList)
	// endregion

	// device
	deviceHandler := handler.NewDeviceHandler()
	/// 获取"获取"描述文件接口
	r.With(middleware.OAuthRegister).Get("/device/config/qr_code", deviceHandler.GetMobileConfigQRCode)
	/// 获取描述文件
	r.With(middleware.OAuthGuest).Get("/device/config/file", deviceHandler.GetMobileConfigFile)
	/// 绑定设备
	r.With(middleware.OAuthGuest).Post("/device/bind/{code}", deviceHandler.Bind)
	// 手动绑定 udid
	r.With(middleware.OAuthRegister).Post("/device/udid", deviceHandler.PostUDID)
	r.With(middleware.OAuthRegister).Post("/device", deviceHandler.PostUDID)
	r.With(middleware.OAuthRegister).Put("/device/{device_id}", deviceHandler.PutUDID)
	r.With(middleware.OAuthRegister).Delete("/device/{device_id}", deviceHandler.DeleteUDID)
	// endregion

	// certificate
	certificateHandler := handler.NewCertificateHandler()
	r.With(middleware.OAuthRegister).Post("/certificate", certificateHandler.Post) /// 生成证书
	r.With(middleware.OAuthRegister).Get("/certificate/price", certificateHandler.GetPrice)
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

	// regis daily_free
	dailyFreeHandler := handler.NewDailyFreeIpaHandler()
	r.With(middleware.OAuthRegister).Post("/daily_free", dailyFreeHandler.PostIpa)
	r.With(middleware.OAuthRegister).Get("/daily_free", dailyFreeHandler.GetDailyFreeRecords)
	// endregion

	// action
	memberActionHandler := handler.NewMemberActionHandler()
	r.With(middleware.OAuthGuest).Get("/member/action", memberActionHandler.GetMemberActions)
	// endregion

	// share_info
	shareInfoHandler := handler.NewShareInfoHandler()
	r.With(middleware.OAuthGuest).Get("/share_info", shareInfoHandler.Get)
	// endregion

	// app_version
	appVersionHandler := handler.NewAppVersionHandler()
	r.With(middleware.OAuthGuest).Get("/app_version", appVersionHandler.CheckAppVersion)
	// endregion

	return r
}

var DefaultRouter = NewRouter()
