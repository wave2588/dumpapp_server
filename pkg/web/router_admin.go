package web

import (
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/handler"
	"dumpapp_server/pkg/web/handler/install_app_handler"
	"github.com/go-chi/chi"
)

func NewRouterAdmin() chi.Router {
	r := chi.NewRouter()

	/// account
	adminAccountHandler := handler.NewAdminAccountHandler()
	r.With(middleware.OAuthAdmin).Post("/account", adminAccountHandler.AddAccount)
	r.With(middleware.OAuthAdmin).Put("/account", adminAccountHandler.PutAccount)
	r.With(middleware.OAuthAdmin).Get("/account/list", adminAccountHandler.AccountList)
	// endregion

	/// config
	adminConfigHandler := handler.NewAdminConfigHandler()
	r.With(middleware.OAuthAdmin).Post("/config", adminConfigHandler.Post)
	r.With(middleware.OAuthAdmin).Get("/config", adminConfigHandler.Get)
	// endregion

	/// admin_member
	adminMemberHandler := handler.NewAdminMemberHandler()
	r.With(middleware.OAuthAdmin).Get("/member", adminMemberHandler.ListMember)
	// endregion

	/// admin_v2
	adminIpaHandler := handler.NewAdminIpaHandler()
	r.With(middleware.OAuthAdmin).Get("/ipa", adminIpaHandler.List)
	r.With(middleware.OAuthAdmin).Post("/ipa", adminIpaHandler.Post)
	r.With(middleware.OAuthAdmin).Get("/ipa/{ipa_id}", adminIpaHandler.Get)
	r.With(middleware.OAuthAdmin).Delete("/ipa", adminIpaHandler.DeleteIpa)
	r.With(middleware.OAuthAdmin).Delete("/batch_ipa", adminIpaHandler.BatchDeleteIpa)
	// endregion

	/// admin 查看未砸壳列表
	adminDumpOrderHandler := handler.NewAdminDumpOrderHandler()
	r.With(middleware.OAuthAdmin).Get("/ipa/dump_order", adminDumpOrderHandler.GetDumpOrderList)
	r.With(middleware.OAuthAdmin).Put("/ipa/dump_order/{dump_order_id}", adminDumpOrderHandler.PutDumpOrderList)
	r.With(middleware.OAuthAdmin).Delete("/ipa/dump_order", adminDumpOrderHandler.DeleteDumpOrder)
	// endregion

	/// admin_record
	adminSearchRecordHandler := handler.NewAdminSearchRecordHandler()
	r.With(middleware.OAuthAdmin).Get("/search/record", adminSearchRecordHandler.GetMemberSearchRecord)
	// endregion

	/// admin download number
	adminMemberPayCountHandler := handler.NewAdminMemberPayCountHandler()
	r.With(middleware.OAuthAdmin).Post("/member/order", adminMemberPayCountHandler.AddNumber)
	r.With(middleware.OAuthAdmin).Delete("/member/order", adminMemberPayCountHandler.DeleteNumber)
	// endregion

	// admin device
	adminDeviceHandler := handler.NewAdminDeviceHandler()
	r.With(middleware.OAuthAdmin).Delete("/device/unbind", adminDeviceHandler.Unbind)
	// endregion

	// admin cdkey handler
	cdkeyHandler := install_app_handler.NewAdminCDKeyHandler()
	r.With(middleware.OAuthAdmin).Post("/cdkey", cdkeyHandler.Post)
	r.With(middleware.OAuthAdmin).Get("/cdkeys", cdkeyHandler.GetList)
	r.With(middleware.OAuthAdmin).Delete("/cdkey/{cdkey_id}", cdkeyHandler.Delete)
	// endregion

	// ipa_black handler
	ipaBlackHandler := handler.NewAdminIpaBlackHandler()
	r.With(middleware.OAuthAdmin).Post("/ipa_black", ipaBlackHandler.Post)
	r.With(middleware.OAuthAdmin).Get("/ipa_black/list", ipaBlackHandler.GetList)
	r.With(middleware.OAuthAdmin).Delete("/ipa_black/{ipa_black_id}", ipaBlackHandler.Delete)
	// endregion

	// dispense handler
	dispenseHandler := handler.NewAdminDispenseHandler()
	r.With(middleware.OAuthAdmin).Post("/dispense", dispenseHandler.AddCount)
	r.With(middleware.OAuthAdmin).Delete("/dispense", dispenseHandler.DeleteCount)
	// endregion

	// auth_website handler
	authWebsiteHandler := handler.NewAdminAuthWebsiteHandler()
	r.With(middleware.OAuthAdmin).Post("/auth/website", authWebsiteHandler.Auth)
	r.With(middleware.OAuthAdmin).Post("/un_auth/website", authWebsiteHandler.UnAuth)
	r.With(middleware.OAuthAdmin).Get("/auth/website/list", authWebsiteHandler.List)
	// endregion

	// certificate handler
	certificateHandler := handler.NewAdminCertificateHandler()
	r.With(middleware.OAuthAdmin).Post("/certificate/replenish", certificateHandler.Replenish)
	// endregion

	return r
}

var DefaultRouterAdmin = NewRouterAdmin()
