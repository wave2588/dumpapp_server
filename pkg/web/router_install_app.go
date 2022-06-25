package web

import (
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/handler/install_app_handler"
	"github.com/go-chi/chi"
)

func NewRouterInstallApp() chi.Router {
	r := chi.NewRouter()

	orderHandler := install_app_handler.NewInstallAppCDKEYOrderHandler()
	r.With(middleware.OAuthGuest).Get("/order", orderHandler.GetOrderURL)

	callbackPayHandler := install_app_handler.NewCallbackPayHandler()
	r.With().Post("/callback_alipay", callbackPayHandler.ALiPayCallback)

	return r
}

var DefaultRouterInstallApp = NewRouterInstallApp()
