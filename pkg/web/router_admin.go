package web

import (
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/handler"
	"github.com/go-chi/chi"
)

func NewRouterAdmin() chi.Router {
	r := chi.NewRouter()

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
	r.With(middleware.OAuthRegister).Post("/ipa", adminIpaHandler.Post)
	r.With(middleware.OAuthRegister).Delete("/ipa", adminIpaHandler.DeleteIpa)
	r.With(middleware.OAuthRegister).Delete("/batch_ipa", adminIpaHandler.BatchDeleteIpa)
	// endregion

	/// admin 查看未砸壳列表
	adminDumpOrderHandler := handler.NewAdminDumpOrderHandler()
	r.With(middleware.OAuthAdmin).Get("/ipa/dump_order", adminDumpOrderHandler.GetDumpOrderList)
	r.With(middleware.OAuthAdmin).Delete("/ipa/dump_order", adminDumpOrderHandler.DeleteDumpOrder)
	// endregion

	/// admin_record
	adminSearchRecordHandler := handler.NewAdminSearchRecordHandler()
	r.With(middleware.OAuthAdmin).Get("/search/record", adminSearchRecordHandler.GetMemberSearchRecord)
	// endregion

	/// admin download order
	adminOrderHandler := handler.NewAdminOrderHandler()
	r.With(middleware.OAuthAdmin).Get("/order", adminOrderHandler.GetOrderCount)
	// endregion

	/// admin download number
	adminDownloadNumberHandler := handler.NewAdminDownloadNumberHandler()
	r.With(middleware.OAuthRegister).Post("/member/download_number", adminDownloadNumberHandler.AddNumber)
	r.With(middleware.OAuthRegister).Delete("/member/download_number", adminDownloadNumberHandler.DeleteNumber)
	// endregion

	return r
}

var DefaultRouterAdmin = NewRouterAdmin()
