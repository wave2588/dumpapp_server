package web

import (
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/handler/install_app_handler"
	"github.com/go-chi/chi"
)

func NewRouterInstallApp() chi.Router {
	r := chi.NewRouter()

	/// order
	orderHandler := install_app_handler.NewCDKEYOrderHandler()
	r.With(middleware.OAuthGuest).Get("/order", orderHandler.GetOrderURL)
	r.With(middleware.OAuthGuest).Get("/order/{order_id}", orderHandler.GetOrderInfo)

	/// cdkey
	cdkeyHandler := install_app_handler.NewCDKEYHandler()
	r.With(middleware.OAuthGuest).Get("/cdkey/{out_id}", cdkeyHandler.GetCDKEYInfo)

	/// certificate
	certificateHandler := install_app_handler.NewCertificateHandler()
	r.With(middleware.OAuthGuest).Post("/certificate", certificateHandler.Post)

	/// ali_pay callback
	callbackPayHandler := install_app_handler.NewCallbackPayHandler()
	r.With().Post("/callback_alipay", callbackPayHandler.ALiPayCallback)

	return r
}

var DefaultRouterInstallApp = NewRouterInstallApp()
