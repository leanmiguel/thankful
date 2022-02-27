package main

import (
	"html/template"
	"net/http"
	"strings"
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

type todayFilledTemplateData struct {
	Entries []string
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

func (app *application) handleTodaySubmission(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	firstEntry := r.Form.Get("first_entry")
	secondEntry := r.Form.Get("second_entry")
	thirdEntry := r.Form.Get("third_entry")

	if firstEntry == "" || secondEntry == "" || thirdEntry == "" {
		http.Error(w, "yeah, you messed up", 400)
	}

	entries := []string{firstEntry, secondEntry, thirdEntry}
	app.entries.Insert("lean", entries)

	http.Redirect(w, r, "/today", http.StatusSeeOther)
}

func (app *application) serveTodayScreen(w http.ResponseWriter, r *http.Request) {

	currentDate := time.Now().UTC().Format(time.RFC3339)

	currentDateParts := strings.Split(currentDate, "T")
	noTimeDate := currentDateParts[0]

	entry, err := app.entries.Get("lean", noTimeDate)

	if err != nil {
		if strings.Contains(err.Error(), "no item found") {
			ts, err := template.ParseFiles("./ui/static/html/today_not_filled.page.tmpl")

			if err != nil {
				app.serverError(w, err)
				http.Error(w, "Internal Server Error", 500)
			}

			err = ts.Execute(w, nil)

			if err != nil {
				app.serverError(w, err)
				http.Error(w, "Internal Server Error", 500)
			}
			return
		} else {
			app.serverError(w, err)
			http.Error(w, "Internal Server Error", 500)
			return
		}
	}

	ts, err := template.ParseFiles("./ui/static/html/today_completed.page.tmpl")

	if err != nil {
		app.serverError(w, err)
		http.Error(w, "internal server error", 500)
	}

	err = ts.Execute(w, todayFilledTemplateData{
		Entries: entry.Entries,
	})

	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
	}

}
func (app *application) serveDayScreen(w http.ResponseWriter, r *http.Request) {

	day := chi.URLParam(r, "day")

	if day == "today" || day == "favicon.ico" {
		return
	}
	ts, err := template.ParseFiles("./ui/static/html/[date].page.tmpl")

	if err != nil {
		app.serverError(w, err)
		http.Error(w, "internal server error", 500)
	}

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
