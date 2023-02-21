package main

import (
	"net/http"

	"github.com/alexedwards/flow"
)

func (app *application) routes() http.Handler {
	mux := flow.New()

	mux.NotFound = http.HandlerFunc(app.notFound)
	mux.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	mux.Use(app.recoverPanic)
	mux.Use(app.authenticate)

	mux.HandleFunc("/status", app.status, "GET")
	mux.HandleFunc("/users", app.createUser, "POST")
	mux.HandleFunc("/authentication-tokens", app.createAuthenticationToken, "POST")

	mux.Group(func(mux *flow.Mux) {
		mux.Use(app.requireAuthenticatedUser)

		mux.HandleFunc("/protected", app.protected, "GET")
	})

	mux.Group(func(mux *flow.Mux) {
		mux.Use(app.requireBasicAuthentication)

		mux.HandleFunc("/basic-auth-protected", app.protected, "GET")
	})

	return mux
}
