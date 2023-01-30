package main

import (
	"gobooking/pkg/config"
	"gobooking/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(ac *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.General)
	mux.Get("/majors-suite", handlers.Repo.Major)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	mux.Get("/search-avaibility", handlers.Repo.SearchAvailbilty)
	mux.Post("/search-avaibility", handlers.Repo.PostSearchAvailbilty)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/favicon.ico", handlers.Repo.DoNothing)

	fileServer := http.FileServer(http.Dir("./public"))

	mux.Handle("/*", http.StripPrefix("/", fileServer))

	return mux
}
