package main

import (
	"fmt"
	"html/template"
	"leanmiguel/thankful/pkg/models"
	"log"
	"net/http"
	"time"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFiles("./ui/static/html/home.page.tmpl")

	if err != nil {
		app.serverError(w, err)
		http.Error(w, "internal server error", 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
	}

}

var mockEntry = models.Entry{
	UserId:      "lean",
	CreatedTime: "2022-02-23T11:53:28Z",
	Entries:     []string{"stuff", "hey", "what"},
}

type templateData struct {
	CreatedTime string
	Entries     []string
}

// var mockEntries = Entries{
// 	mockEntry,
// }
const layoutUS = "January 2, 2006"

func (app *application) serveTodayScreen(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFiles("./ui/static/html/[date].page.tmpl")

	if err != nil {

		log.Println(err.Error())
		http.Error(w, "internal server error", 500)
	}

	// fmt.Println(time.Now().UTC().Format(time.RFC3339))
	createdTime, err := time.Parse(time.RFC3339, mockEntry.CreatedTime)
	formattedTime := createdTime.Format(layoutUS)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Time parse no bueno", 500)

	}

	err = ts.Execute(w, templateData{
		CreatedTime: formattedTime,
		Entries:     mockEntry.Entries,
	})
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}
func (app *application) serveDayScreen(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"name":"not alex"}`))
}
