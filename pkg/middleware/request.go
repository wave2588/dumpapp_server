package middleware

import (
	"context"
	"net/http"

	"dumpapp_server/pkg/common/constant"
)

func RequestContext(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		rContext := context.WithValue(r.Context(), constant.CtxKeyAppVersion, r.Header.Get("x-app-version"))
		rContext = context.WithValue(rContext, constant.CtxKeyAppPlatform, r.Header.Get("x-app-platform"))
		rContext = context.WithValue(rContext, constant.CtxKeyAppUDID, r.Header.Get("x-app-udid"))
		h.ServeHTTP(w, r.WithContext(rContext))
	}
	return http.HandlerFunc(fn)
}
