package main

import (
	"net/http"

	"github.com/sakaguchi-0725/bookings/pkg/config"
	"github.com/sakaguchi-0725/bookings/pkg/handlers"

	"github.com/go-chi/chi/v4"
	"github.com/go-chi/chi/v4/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileSerer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileSerer))
	return mux
}
