package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *Server) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	mux.Handle("/calendars/", http.StripPrefix("/calendars", app.calendarRoutes()))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}

// RegisterCalendarRoutes registers routes for the calendar resource(s)
//
// @param app *Server - the server instance that implements the handlers
//
// @return *http.ServeMux - the router with the calendar routes registered
func (app *Server) calendarRoutes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /", app.getCalendars)

	r.HandleFunc("GET /{id}", app.getCalendar)

	r.HandleFunc("POST /", app.createCalendar)

	r.HandleFunc("PUT /{id}", app.updateCalendar)

	r.HandleFunc("DELETE /{id}", app.deleteCalendar)

	return r
}
