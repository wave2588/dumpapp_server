package middleware

import (
	"net/http"
	"strings"
	"sync"

	"dumpapp_server/pkg/common/constant"
	"github.com/go-chi/cors"
)

/// 解决跨域问题
func Cors(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Expose-Headers", "Set-Cookie")
		// 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") // header的类型
		w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    // 设置为true，允许ajax异步请求带cookie信息
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             // 允许请求方法
		w.Header().Set("content-type", "application/json;charset=UTF-8")                                              // 返回数据格式是json
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func CORS(origins []string) func(http.Handler) http.Handler {
	defaultCORS := cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", constant.AppOpsAuthNameHeaderKey},
		AllowCredentials: true,
		MaxAge:           3600,
	})
	corsMapPool := &sync.Map{}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			headers := r.Header.Get("Access-Control-Request-Headers")
			if headers != "" {
				var p *sync.Pool
				if v, ok := corsMapPool.Load(headers); ok {
					p = v.(*sync.Pool)
				} else {
					opt := cors.Options{
						AllowedOrigins:   origins,
						AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
						AllowedHeaders:   strings.Split(headers, ","),
						AllowCredentials: true,
						MaxAge:           3600,
					}
					for i, header := range opt.AllowedHeaders {
						opt.AllowedHeaders[i] = strings.TrimSpace(header)
					}

					p = &sync.Pool{
						New: func() interface{} {
							return cors.New(opt)
						},
					}
					corsMapPool.Store(headers, p)
				}
				cc := p.Get().(*cors.Cors)
				cc.Handler(next).ServeHTTP(w, r)
				p.Put(cc)
			} else {
				defaultCORS.Handler(next).ServeHTTP(w, r)
			}
		})
	}
}
