package web

import (
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/handler/open_api_handler"
	"github.com/go-chi/chi"
)

func NewRouterOpenAPI() chi.Router {
	r := chi.NewRouter()

	certificateHandler := open_api_handler.NewOpenCertificateHandler()
	r.With(middleware.OAuthGuest).Post("/certificate", certificateHandler.PostCertificate)
	r.With(middleware.OAuthGuest).Get("/certificate", certificateHandler.GetCertificate)
	r.With(middleware.OAuthGuest).Get("/certificate/list", certificateHandler.GetCertificateList)
	r.With(middleware.OAuthGuest).Get("/certificate/price", certificateHandler.GetCertificatePrice)

	return r
}

var DefaultRouterOpenAPI = NewRouterOpenAPI()
