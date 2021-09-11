package web

import (
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/handler"
	"github.com/go-chi/chi"
)

func NewRouterAdmin() chi.Router {
	r := chi.NewRouter()

	/// admin_member
	adminMemberHandler := handler.NewAdminMemberHandler()
	r.With(middleware.OAuthAdmin).Get("/member", adminMemberHandler.ListMember)
	// endregion

	/// admin_v2
	adminIpaHandler := handler.NewAdminIpaHandler()
	r.With(middleware.OAuthRegister).Post("/ipa", adminIpaHandler.Post)
	r.With(middleware.OAuthRegister).Delete("/ipa", adminIpaHandler.DeleteIpa)
	r.With(middleware.OAuthRegister).Delete("/batch_ipa", adminIpaHandler.BatchDeleteIpa)
	// endregion

	/// admin_record
	adminSearchRecordHandler := handler.NewAdminSearchRecordHandler()
	r.With(middleware.OAuthAdmin).Get("/search/record", adminSearchRecordHandler.GetMemberSearchRecord)
	// endregion

	/// admin download number
	adminDownloadNumberHandler := handler.NewAdminDownloadNumberHandler()
	r.With(middleware.OAuthRegister).Post("/member/download_number", adminDownloadNumberHandler.AddNumber)
	r.With(middleware.OAuthRegister).Delete("/member/download_number", adminDownloadNumberHandler.DeleteNumber)
	// endregion

	return r
}

var DefaultRouterAdmin = NewRouterAdmin()
