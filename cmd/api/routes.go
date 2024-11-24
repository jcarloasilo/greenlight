package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)

	mux.Route("/v1", func(r chi.Router) {

		r.Get("/healthcheck", app.healthcheckHandler)

		r.Post("/movies", app.createMovieHandler)
		r.Get("/movies/{id}", app.showMovieHandler)

	})

	return mux
}