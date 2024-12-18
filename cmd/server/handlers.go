package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// createCalendar creates a new calendar
func (app *Server) createCalendar(w http.ResponseWriter, req *http.Request) {
	var input struct {
		Name    *string `json:"name"`
		OwnerID *int    `json:"owner_id"`
	}

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		app.logger.Errorw("Failed to decode request body", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if input.Name == nil || *input.Name == "" || input.OwnerID == nil {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	cal, err := app.calendars.Create(*input.Name, *input.OwnerID)
	if err != nil {
		app.logger.Errorw("Failed to save calendar", err)
		http.Error(w, "Failed to create calendar", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cal); err != nil {
		app.logger.Errorw("Failed to encode response", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// getCalendar retrieves a calendar by ID
func (app *Server) getCalendar(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.NotFound(w, req)
		return
	}

	cal, err := app.calendars.Get(id)
	if err != nil {
		app.logger.Errorw("Failed to get calendar", err)
		http.Error(w, "Failed to get calendar", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cal); err != nil {
		app.logger.Errorw("Failed to encode response", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// getCalendars retrieves all calendars
func (app *Server) getCalendars(w http.ResponseWriter, req *http.Request) {
	cals, err := app.calendars.GetAll()
	if err != nil {
		app.logger.Errorw("Failed to get calendars", err)
		http.Error(w, "Failed to get calendars", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cals); err != nil {
		app.logger.Errorw("Failed to encode response", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// updateCalendar updates a calendar by ID
func (app *Server) updateCalendar(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.NotFound(w, req)
		return
	}

	var input struct {
		Name *string `json:"name"`
	}

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		app.logger.Errorw("Failed to decode request body", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if input.Name == nil || *input.Name == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	cal, err := app.calendars.Update(id, *input.Name)
	if err != nil {
		app.logger.Errorw("Failed to update calendar", err)
		http.Error(w, "Failed to update calendar", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cal); err != nil {
		app.logger.Errorw("Failed to encode response", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (app *Server) deleteCalendar(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.NotFound(w, req)
		return
	}

	if err := app.calendars.Delete(id); err != nil {
		app.logger.Errorw("Failed to delete calendar", err)
		http.Error(w, "Failed to delete calendar", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
