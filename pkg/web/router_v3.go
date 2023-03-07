package web

import (
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/handler"
	"github.com/go-chi/chi"
)

func NewRouterV3() chi.Router {
	r := chi.NewRouter()

	// member pay order
	memberPayOrder := handler.NewMemberPayOrderHandler()
	r.With(middleware.OAuthRegister).Get("/member/order", memberPayOrder.GetPayOrderURL)
	r.With(middleware.OAuthRegister).Get("/member/order/{order_id}", memberPayOrder.GetOrder)
	r.With(middleware.OAuthRegister).Get("/member/order/rule", memberPayOrder.GetOrderRule)

	r.With().Get("/member/order/test", memberPayOrder.GetPayOrderURLTest)

	// ipa
	downloadHandler := handler.NewDownloadHandler()
	r.With(middleware.OAuthRegister).Get("/ipa/{ipa_id}/download_url", downloadHandler.GetDownloadURL)
	r.With(middleware.OAuthRegister).Get("/ipa/{ipa_id}/check_can_download", downloadHandler.CheckCanDownload)

	/// callback_pay
	callbackPayV3Handler := handler.NewCallbackPayV3Handler()
	r.With().Post("/callback_alipay", callbackPayV3Handler.ALiPayCallback)

	return r
}

var DefaultRouterV3 = NewRouterV3()
