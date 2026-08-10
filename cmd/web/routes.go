package main

import (
	"net/http"

	"snipbox/internal/chains"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("POST /snippet/create", app.snippetCreate)

	standard := chains.New(app.recoverPanic, app.logRequest, secureHearders)

	return standard.Then(mux)
}
