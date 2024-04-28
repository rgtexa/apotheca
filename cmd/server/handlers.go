package main

import "net/http"

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}
