package commons

import (
	"log/slog"
)

type PingableDB interface {
	Ping() error
}

type DbHealthIndicator struct {
	Name   string
	Db     PingableDB
	Logger *slog.Logger
}

func NewDbHealthIndicator(name string, db PingableDB, logger *slog.Logger) *DbHealthIndicator {
	return &DbHealthIndicator{
		Name:   name,
		Db:     db,
		Logger: logger,
	}
}

func (dbHealthIndicator *DbHealthIndicator) IndicateHealth() (string, bool) {
	err := dbHealthIndicator.Db.Ping()
	if err != nil {
		dbHealthIndicator.Logger.Error("Error connecting to database", slog.String("error", err.Error()))
		return dbHealthIndicator.Name, false
	}
	return dbHealthIndicator.Name, true
}
