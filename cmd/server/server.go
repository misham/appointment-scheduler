package main

import (
	"log"

	"github.com/misham/appointment-scheduler/pkg/calendar"
	"go.uber.org/zap"
)

// Server handles dependency injection for route handlers
type Server struct {
	logger    *zap.SugaredLogger
	errorLog  *log.Logger
	calendars calendar.CalendarModelInterface
}
