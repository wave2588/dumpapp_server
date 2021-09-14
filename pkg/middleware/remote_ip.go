package middleware

import (
	"context"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/middleware/util"
)

func RemoteIP(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ip := util.GetRemoteIP(r)
		ctx := context.WithValue(r.Context(), constant.RemoteIP, ip)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
