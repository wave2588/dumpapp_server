package middleware

import (
	"context"
	"dumpapp_server/pkg/common/constant"
	"net/http"
)

func RequestContext(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		rContext := context.WithValue(r.Context(), constant.CtxKeyAppVersion, r.Header.Get("x-app-version"))
		rContext = context.WithValue(rContext, constant.CtxKeyAppPlatform, r.Header.Get("x-app-platform"))
		h.ServeHTTP(w, r.WithContext(rContext))
	}
	return http.HandlerFunc(fn)
}
