package main

import (
	"dumpapp_server/pkg/common/util"
	middleware2 "dumpapp_server/pkg/middleware"
	"dumpapp_server/pkg/web"

	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	fmt.Println("start run")
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.NoCache)
	r.Use(middleware.Heartbeat("/"))
	r.Use(middleware.Heartbeat("/check_health"))

	r.Use(middleware2.PanicAsError)
	//r.Use(middleware2.Cors)
	r.Use(middleware2.CORS(origins()))

	r.Mount("/api/", web.DefaultRouter)
	r.Mount("/api/v2/", web.DefaultRouterV2)

	util.PanicIf(http.ListenAndServe(":1995", r))
}

func origins() []string {
	return []string{
		"https://dumpapp.com",
		"https://www.dumpapp.com",
		"http://dumpapp.com",
		"http://www.dumpapp.com",
		"http://127.0.0.1:8080",
		"http://localhost:8080",
	}
}
