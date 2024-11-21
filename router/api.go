package router

import (
	"e-commerce-app/infra"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupReouter(ctx infra.Context) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/customer", func(r chi.Router) {
		r.Post("/", ctx.Handler.CustomerHandler.Create)
		r.Get("/", ctx.Handler.CustomerHandler.GetAll)
		r.Get("/{id}", ctx.Handler.CustomerHandler.Create)
	})

	fmt.Println("server start on port ", ctx.Config.Port)
	http.ListenAndServe(":"+ctx.Config.Port, r)
}
