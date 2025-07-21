package commons

import (
	"log/slog"
)

// PingableDB is a DB that can be pinged.
type PingableDB interface {
	Ping() error
}

// DbHealthIndicator is a HealthIndicator for a given PingableDB.
type DbHealthIndicator struct {
	Name   string
	Db     PingableDB
	Logger *slog.Logger
}

// NewDbHealthIndicator is a factory method for producing a DbHealthIndicator.
func NewDbHealthIndicator(name string, db PingableDB, logger *slog.Logger) *DbHealthIndicator {
	return &DbHealthIndicator{
		Name:   name,
		Db:     db,
		Logger: logger,
	}
}

// IndicateHealth is the HealthIndicator.IndicateHealth implementation for DbHealthIndicator.
func (dbHealthIndicator *DbHealthIndicator) IndicateHealth() (string, bool) {
	err := dbHealthIndicator.Db.Ping()
	if err != nil {
		dbHealthIndicator.Logger.Error("Error connecting to database", slog.String("error", err.Error()))

		return dbHealthIndicator.Name, false
	}

	return dbHealthIndicator.Name, true
}
