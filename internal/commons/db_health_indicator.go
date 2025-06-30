package commons

import (
	"database/sql"
	"log/slog"
)

type DbHealthIndicator struct {
	Name   string
	Db     *sql.DB
	Logger *slog.Logger
}

func NewDbHealthIndicator(name string, db *sql.DB, logger *slog.Logger) *DbHealthIndicator {
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
