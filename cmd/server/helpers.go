package server

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/justinas/nosurf"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, td *templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		app.serverError(w, r, fmt.Errorf("the template %s could not be found", page))
		return
	}
	td.ReturnURI = r.URL.RequestURI()

	buf := new(bytes.Buffer)
	err := ts.ExecuteTemplate(buf, "base", td)
	if err != nil {
		app.serverError(w, r, err)
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
		ReturnURI:       "",
	}
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
