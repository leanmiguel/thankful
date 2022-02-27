package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {

	r := chi.NewRouter()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	r.Get("/", app.home)
	r.Get("/today", app.serveTodayScreen)
	r.Post("/today", app.handleTodaySubmission)
	r.Get("/{day}", app.serveDayScreen)
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return r
}
