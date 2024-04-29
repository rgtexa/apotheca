package main

import (
	"bytes"
	"html/template"
	"log/slog"
	"net/http"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./ui/html/base.html")
	if err != nil {
		app.logger.Error("failed to parse template", slog.String("error", err.Error()))
	}
	buf := new(bytes.Buffer)
	err = tmp.ExecuteTemplate(buf, "base.html", nil)
	if err != nil {
		app.logger.Error("failed to execute template", slog.String("error", err.Error()))
		return
	}

	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusOK)

	buf.WriteTo(w)
}
