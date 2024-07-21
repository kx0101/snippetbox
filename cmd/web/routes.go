package main

import (
	"net/http"
)

func (app *application) route() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.Handle("GET /{$}", app.sessionManager.LoadAndSave(http.HandlerFunc(app.home)))
	mux.Handle("GET /snippet/view/{id}", app.sessionManager.LoadAndSave(http.HandlerFunc(app.snippetView)))
	mux.Handle("GET /snippet/create", app.sessionManager.LoadAndSave(http.HandlerFunc(app.snippetCreate)))
	mux.Handle("POST /snippet/create", app.sessionManager.LoadAndSave(http.HandlerFunc(app.snippetCreatePost)))

	mux.Handle("GET /user/signup", app.sessionManager.LoadAndSave((http.HandlerFunc(app.userSignup))))
	mux.Handle("POST /user/signup", app.sessionManager.LoadAndSave((http.HandlerFunc(app.userSignupPost))))
	mux.Handle("GET /user/login", app.sessionManager.LoadAndSave((http.HandlerFunc(app.userLogin))))
	mux.Handle("POST /user/login", app.sessionManager.LoadAndSave((http.HandlerFunc(app.userLoginPost))))
	mux.Handle("POST /user/logout", app.sessionManager.LoadAndSave((http.HandlerFunc(app.userLogoutPost))))

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
