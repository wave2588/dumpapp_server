package web

import (
	"dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web/handler"
	"github.com/go-chi/chi"
)

func NewRouterV2() chi.Router {
	r := chi.NewRouter()

	// member_vip
	memberVipV2Handler := handler.NewMemberVipV2Handler()
	r.With(middleware.OAuthRegister).Get("/member/vip", memberVipV2Handler.GetV2)
	r.With(middleware.OAuthRegister).Post("/member/vip", memberVipV2Handler.GetPayURL)

	// ipa
	downloadHandler := handler.NewDownloadHandler()
	r.With(middleware.OAuthRegister).Get("/ipa/{ipa_id}/download_url", downloadHandler.GetDownloadURL)
	r.With(middleware.OAuthRegister).Get("/ipa/{ipa_id}/check_can_download", downloadHandler.CheckCanDownload)

	/// callback_pay
	callbackPayV2Handler := handler.NewCallbackPayV2Handler()
	r.With().Post("/callback_alipay", callbackPayV2Handler.ALiPayCallback)

	return r
}

var DefaultRouterV2 = NewRouterV2()
