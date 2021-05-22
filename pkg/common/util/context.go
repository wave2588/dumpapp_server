package util

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

// URLParam returns the url parameter from a http.Request object.
func URLParam(r *http.Request, key string) string {
	if rctx := RouteContext(r.Context()); rctx != nil {
		return rctx.URLParam(key)
	}
	return ""
}

// RouteContext returns chi's routing Context object from a
// http.Request Context.
func RouteContext(ctx context.Context) *chi.Context {
	val, _ := ctx.Value(chi.RouteCtxKey).(*chi.Context)
	return val
}
