package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer) //如果炸了不會讓你整個app炸回home會有backup
	mux.Use(app.enableCORS)
	mux.Get("/", app.Home)

	mux.Get("/movies", app.AllMovies)

	return mux
}