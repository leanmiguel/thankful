package main

import (
	"html/template"
	"net/http"
	"time"
)

const layoutUS = "January 2, 2006"

// entries := &[]models.Entry{
// 	{
// 		UserId:      "lean",
// 		CreatedTime: "2022-02-22T11:53:28Z",
// 		Entries:     []string{"1", "2", "3"},
// 	},
// 	{
// 		UserId:      "lean",
// 		CreatedTime: "2022-02-23T11:53:28Z",
// 		Entries:     []string{"1", "2", "3"},
// 	},
// }

type EntryData struct {
	Date    string
	DateURL template.URL
}
type homeTemplateData struct {
	AvailableDates []EntryData
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
	app.entries.Get("lean", "2022-02-22")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"name":"not alex"}`))
}
