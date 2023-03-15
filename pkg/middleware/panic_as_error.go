package middleware

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/sentry"
	"dumpapp_server/pkg/errors"
	"dumpapp_server/pkg/web/formatter"
	pkgErrors "github.com/pkg/errors"
)

func PanicAsError(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			recoverd := recover()
			if recoverd == nil {
				return
			}

			if err, ok := recoverd.(error); ok {
				/// 可以介入 sentry
				switch realErr := pkgErrors.Cause(err).(type) {
				case *errors.APIError:
					formatter.RenderError(w, realErr)
					return
				default:
					fmt.Println(err)
					sentry.RavenCaptureError(err)
					formatter.RenderError(w, errors.NewDefaultAPIError(http.StatusInternalServerError, 50000, "InternalServerError", "服务器错误"))
					return
				}
			} else {
				panic(recoverd)
			}
		}()
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
