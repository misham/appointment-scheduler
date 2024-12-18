package main

import (
	"net/http"
	"runtime/debug"
)

// The serverError logs the error and sends a generic 500 response.
func (app *Server) serverError(w http.ResponseWriter, r *http.Request, err error) {
	method := r.Method
	uri := r.URL.RequestURI()
	trace := string(debug.Stack()) // TODO only if DEBUG is set???

	app.logger.Errorw(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError sends specified status code to the client.
func (app *Server) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
