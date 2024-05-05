package server

import (
	"bytes"
	"log/slog"
	"net/http"

	"github.com/justinas/nosurf"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, td *templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		app.logger.Error("template does not exist", slog.String("page", page))
		return
	}

	buf := new(bytes.Buffer)
	err := ts.ExecuteTemplate(buf, "base", td)
	if err != nil {
		app.logger.Error("failed to execute template", slog.String("error", err.Error()))
		return
	}
	w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		Flash:           "",
		IsAuthenticated: false,
		CSRFToken:       nosurf.Token(r),
	}
}
