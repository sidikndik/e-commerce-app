package router

import (
	"e-commerce-app/infra"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupReouter(ctx infra.Context) {
	r := chi.NewRouter()

	r.Route("/customer ", func(r chi.Router) {
		r.Post("/", ctx.Handler.CustomerHandler.Create)
	})

	fmt.Println("server start on port ", ctx.Config.Port)
	http.ListenAndServe(":"+ctx.Config.Port, r)
}
