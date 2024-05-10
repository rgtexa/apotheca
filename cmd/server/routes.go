package server

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	dynamic := alice.New(noSurf, app.sessionManager.LoadAndSave)

	files := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static", http.NotFoundHandler())
	mux.Handle("GET /static/", http.StripPrefix("/static", files))

	mux.Handle("GET /", dynamic.ThenFunc(app.homeHandler))
	mux.Handle("GET /about", dynamic.ThenFunc(app.aboutHandler))
	mux.Handle("GET /license", dynamic.ThenFunc(app.licenseHandler))
	mux.Handle("GET /privacy", dynamic.ThenFunc(app.privacyHandler))
	mux.Handle("GET /contact", dynamic.ThenFunc(app.contactHandler))

	std := alice.New(app.recoverPanic, commonHeaders)

	return std.Then(mux)
}
