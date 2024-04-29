package main

import (
	"net/http"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	w.Header().Set("Content-Type", "text/html")

	app.render(w, r, http.StatusOK, "home.html", data)
}

func (app *application) aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, r, http.StatusOK, "about.html", data)
}

func (app *application) privacyHandler(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, r, http.StatusOK, "privacy.html", data)
}

func (app *application) licenseHandler(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, r, http.StatusOK, "license.html", data)
}

func (app *application) contactHandler(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, r, http.StatusOK, "contact.html", data)
}
