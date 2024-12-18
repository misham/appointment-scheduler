// Package calendar implements the business logic for managing calendars.
package calendar

import "fmt"

// CalendarModelInterface defines the methods for interacting with the calendar model
type CalendarModelInterface interface {
	Create(string, int) (*Calendar, error)
	Get(int) (*Calendar, error)
	GetAll() ([]*Calendar, error)
	Update(int, string) (*Calendar, error)
	Delete(int) error
}

// Calendar represents a calendar entity
type Calendar struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	OwnerID int    `json:"owner_id"`
}

type CalendarModel struct {
	store map[int]*Calendar
}

// NewCalendarModel creates a new CalendarModel instance
//
// @return *CalendarModel - the new CalendarModel instance
func NewCalendarModel() *CalendarModel {
	return &CalendarModel{
		store: make(map[int]*Calendar),
	}
}

// Create a new calendar
//
// @param name string - the name of the calendar
// @param ownerID int - the ID of the owner of the calendar
//
// @return *Calendar - the created calendar entity
// @return error - an error if the creation fails
func (m *CalendarModel) Create(name string, ownerID int) (*Calendar, error) {
	cal := &Calendar{
		ID:      len(m.store) + 1,
		Name:    name,
		OwnerID: ownerID,
	}
	m.store[len(m.store)+1] = cal

	return cal, nil
}

// Get a calendar by ID
//
// @param id int - the ID of the calendar to retrieve
//
// @return *Calendar - the calendar entity
// @return error - an error if the retrieval fails
func (m *CalendarModel) Get(id int) (*Calendar, error) {
	calendar, exists := m.store[id]
	if !exists {
		return nil, fmt.Errorf("calendar with id %d not found", id)
	}
	return calendar, nil
}

// GetAll retrieves all calendars
//
// @return []*Calendar - a list of all calendars
// @return error - an error if the retrieval fails
func (m *CalendarModel) GetAll() ([]*Calendar, error) {
	cals := make([]*Calendar, 0, len(m.store))

	for _, cal := range m.store {
		cals = append(cals, cal)
	}

	return cals, nil
}

// Update a calendar by ID
//
// @param id int - the ID of the calendar to update
// @param name string - the new name of the calendar
//
// @return error - an error if the update fails
func (m *CalendarModel) Update(id int, name string) (*Calendar, error) {
	cal, exists := m.store[id]
	if !exists {
		return nil, fmt.Errorf("calendar with id %d not found", id)
	}

	cal.Name = name
	m.store[id] = cal

	return cal, nil
}

// Delete a calendar by ID
//
// @param id int - the ID of the calendar to delete
//
// @return error - an error if the deletion fails
func (m *CalendarModel) Delete(id int) error {
	_, exists := m.store[id]
	if !exists {
		return fmt.Errorf("calendar with id %d not found", id)
	}

	delete(m.store, id)

	return nil
}
