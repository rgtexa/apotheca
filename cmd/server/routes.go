package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	dynamic := alice.New(noSurf)

	files := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static", http.NotFoundHandler())
	mux.Handle("GET /static/", http.StripPrefix("/static", files))

	mux.Handle("GET /", dynamic.ThenFunc(app.homeHandler))

	std := alice.New(commonHeaders)

	return std.Then(mux)
}
