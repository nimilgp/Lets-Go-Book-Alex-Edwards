package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) route(cfg config) http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(cfg.staticDir))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.getRoot))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.getSnippetView))
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(app.getSnippetCreate))
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.postSnippetCreate))

	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.getUserSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.postUserSignup))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.getUserLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.postUserLogin))
	mux.Handle("POST /user/logout", dynamic.ThenFunc(app.postUserLogout))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
