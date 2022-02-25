package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

const layoutUS = "January 2, 2006"

type EntryData struct {
	Date    string
	DateURL template.URL
}
type homeTemplateData struct {
	AvailableDates []EntryData
}

type dayTemplateData struct {
	Entries     []string
	CreatedDate string
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFiles("./ui/static/html/home.page.tmpl")

	if err != nil {
		app.serverError(w, err)
		http.Error(w, "internal server error", 500)
	}

	entries, err := app.entries.Latest("lean")

	if err != nil {
		app.serverError(w, err)
	}

	var availableDates []EntryData

	for _, entry := range *entries {

		createdTime, err := time.Parse(time.RFC3339, entry.CreatedTime)

		if err != nil {
			app.serverError(w, err)
		}

		formattedTime := createdTime.Format(layoutUS)

		availableDates = append(availableDates, EntryData{Date: formattedTime, DateURL: template.URL(entry.CreatedTime)})
	}

	err = ts.Execute(w, homeTemplateData{
		AvailableDates: availableDates,
	})

	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
	}

}

func (app *application) serveTodayScreen(w http.ResponseWriter, r *http.Request) {

	// ts, err := template.ParseFiles("./ui/static/html/[date].page.tmpl")

	// if err != nil {

	// 	log.Println(err.Error())
	// 	http.Error(w, "internal server error", 500)
	// }

	// createdTime, err := time.Parse(time.RFC3339, mockEntry.CreatedTime)
	// formattedTime := createdTime.Format(layoutUS)

	// if err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, "Time parse no bueno", 500)

	// }

	// err = ts.Execute(w, templateData{
	// 	CreatedTime: formattedTime,
	// 	Entries:     mockEntry.Entries,
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// 	http.Error(w, "Internal Server Error", 500)
	// }

}
func (app *application) serveDayScreen(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFiles("./ui/static/html/[date].page.tmpl")

	if err != nil {
		app.serverError(w, err)
		http.Error(w, "internal server error", 500)
	}

	day := chi.URLParam(r, "day")

	entry, err := app.entries.Get("lean", day)

	if err != nil {
		app.serverError(w, err)
		http.Error(w, "internal server error", 500)
	}

	createdTime, err := time.Parse(time.RFC3339, entry.CreatedTime)
	formattedTime := createdTime.Format(layoutUS)

	if err != nil {
		app.serverError(w, err)
	}

	err = ts.Execute(w, dayTemplateData{
		Entries:     entry.Entries,
		CreatedDate: formattedTime,
	})

	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
	}
}
