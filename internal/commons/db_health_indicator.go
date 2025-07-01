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

func (appDatabase *DbHealthIndicator) IndicateHealth() (string, bool) {
	err := appDatabase.Db.Ping()
	if err != nil {
		appDatabase.Logger.Error("Error connecting to database", slog.String("error", err.Error()))
		return appDatabase.Name, false
	}
	return appDatabase.Name, true
}
