package handlers

import (
	"github.com/go-chi/chi"

	"github.com/Stransyyy/GoRestApi/internal/middleware"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(r chi.Router) {

		r.Use(middleware.Authorization)

		r.Get("/task", GetTaskHandler)
	})
}
