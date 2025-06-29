package commons

import (
	"database/sql"
	"fmt"
	"log/slog"
)

type AppDatabase struct {
	Name   string
	Db     *sql.DB
	Logger *slog.Logger
}

func NewAppDatabase(name string, driver string, dataSource string, logger *slog.Logger) (*AppDatabase, error) {
	database, err := sql.Open(driver, dataSource)
	if err != nil {
		logger.Error(
			fmt.Sprintf("Failed to open database: %s", name),
			slog.String("driver", driver),
			slog.String("datasource", dataSource),
			slog.String("error", err.Error()))
		return nil, err
	}

	return &AppDatabase{
		Name:   name,
		Db:     database,
		Logger: logger,
	}, nil
}

func (appDatabase *AppDatabase) IndicateHealth() (string, bool) {
	err := appDatabase.Db.Ping()
	if err != nil {
		appDatabase.Logger.Error("Error connecting to database", slog.String("error", err.Error()))
		return appDatabase.Name, false
	}
	return appDatabase.Name, true
}
