package router

import (
	"github.com/go-chi/chi/v5"
)

func SetupReouter() *chi.Mux {
	r := chi.NewRouter()

	return r
}
