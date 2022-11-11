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
	r.With(middleware.OAuthAdminV2).Post("/account", adminAccountHandler.AddAccount)
	r.With(middleware.OAuthAdminV2).Put("/account", adminAccountHandler.PutAccount)
	r.With(middleware.OAuthAdmin).Get("/account/list", adminAccountHandler.AccountList)
	r.With(middleware.OAuthAdmin).Get("/account", adminAccountHandler.GetAccount)
	// endregion

	/// config
	adminConfigHandler := handler.NewAdminConfigHandler()
	r.With(middleware.OAuthAdmin).Post("/config", adminConfigHandler.Post)
	r.With(middleware.OAuthAdminV2).Get("/config", adminConfigHandler.Get)
	// endregion

	/// admin_member
	adminMemberHandler := handler.NewAdminMemberHandler()
	r.With(middleware.OAuthAdmin).Get("/member", adminMemberHandler.ListMember)
	// endregion

	/// admin_v2
	adminIpaHandler := handler.NewAdminIpaHandler()
	r.With(middleware.OAuthAdminV2).Get("/ipa", adminIpaHandler.List)
	r.With(middleware.OAuthAdminV2).Post("/ipa", adminIpaHandler.Post)
	r.With(middleware.OAuthAdminV2).Get("/ipa/{ipa_id}", adminIpaHandler.Get)
	r.With(middleware.OAuthAdmin).Delete("/ipa", adminIpaHandler.DeleteIpa)
	r.With(middleware.OAuthAdmin).Delete("/batch_ipa", adminIpaHandler.BatchDeleteIpa)
	// endregion

	/// admin 查看未砸壳列表
	adminDumpOrderHandler := handler.NewAdminDumpOrderHandler()
	r.With(middleware.OAuthAdminV2).Get("/ipa/dump_order", adminDumpOrderHandler.GetDumpOrderList)
	r.With(middleware.OAuthAdminV2).Put("/ipa/dump_order/{dump_order_id}", adminDumpOrderHandler.PutDumpOrderList)
	r.With(middleware.OAuthAdminV2).Delete("/ipa/dump_order", adminDumpOrderHandler.DeleteDumpOrder)
	// endregion

	/// admin_record
	adminSearchRecordHandler := handler.NewAdminSearchRecordHandler()
	r.With(middleware.OAuthAdmin).Get("/search/record", adminSearchRecordHandler.GetMemberSearchRecord)
	// endregion

	/// admin download number
	adminMemberPayCountHandler := handler.NewAdminMemberPayCountHandler()
	r.With(middleware.OAuthAdminV2).Post("/member/order", adminMemberPayCountHandler.AddNumber)
	r.With(middleware.OAuthAdminV2).Delete("/member/order", adminMemberPayCountHandler.DeleteNumber)
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
	r.With(middleware.OAuthAdminV2).Post("/ipa_black", ipaBlackHandler.Post)
	r.With(middleware.OAuthAdminV2).Get("/ipa_black/list", ipaBlackHandler.GetList)
	r.With(middleware.OAuthAdminV2).Delete("/ipa_black/{ipa_black_id}", ipaBlackHandler.Delete)
	// endregion

	// dispense handler
	dispenseHandler := handler.NewAdminDispenseHandler()
	r.With(middleware.OAuthAdmin).Post("/dispense", dispenseHandler.AddCount)
	r.With(middleware.OAuthAdmin).Delete("/dispense", dispenseHandler.DeleteCount)
	// endregion

	// auth_website handler
	authWebsiteHandler := handler.NewAdminAuthWebsiteHandler()
	r.With(middleware.OAuthAdminV2).Post("/auth/website", authWebsiteHandler.Auth)
	r.With(middleware.OAuthAdminV2).Post("/un_auth/website", authWebsiteHandler.UnAuth)
	r.With(middleware.OAuthAdminV2).Get("/auth/website/list", authWebsiteHandler.List)
	// endregion

	// certificate handler
	certificateHandler := handler.NewAdminCertificateHandler()
	r.With(middleware.OAuthAdmin).Post("/certificate/replenish", certificateHandler.Replenish)
	// endregion

	return r
}

var DefaultRouterAdmin = NewRouterAdmin()
