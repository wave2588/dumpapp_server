package web

import (
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/handler/open_api_handler"
	"github.com/go-chi/chi"
)

func NewRouterOpenAPI() chi.Router {
	r := chi.NewRouter()

	// ipa handler
	ipaHandler := open_api_handler.NewOpenIpaHandler()
	r.With(middleware.OAuthGuest).Get("/ipa", ipaHandler.Get)
	r.With(middleware.OAuthGuest).Get("/ipa/download_url", ipaHandler.GetIpaDownloadURL)
	// endregion

	// certificate handler
	certificateHandler := open_api_handler.NewOpenCertificateHandler()
	r.With(middleware.OAuthGuest).Post("/certificate", certificateHandler.PostCertificate)
	r.With(middleware.OAuthGuest).Get("/certificate", certificateHandler.GetCertificate)
	r.With(middleware.OAuthGuest).Get("/certificate/list", certificateHandler.GetCertificateList)
	r.With(middleware.OAuthGuest).Get("/certificate/price", certificateHandler.GetCertificatePrice)
	// endregion

	// member handler
	memberHandler := open_api_handler.NewOpenMemberHandler()
	r.With(middleware.OAuthGuest).Get("/member", memberHandler.GetMember)
	// endregion

	// access website
	authWebsiteHandler := open_api_handler.NewAuthWebsiteHandler()
	r.With(middleware.OAuthGuest).Get("/auth/website", authWebsiteHandler.GetAuth)
	// endregion

	return r
}

var DefaultRouterOpenAPI = NewRouterOpenAPI()
