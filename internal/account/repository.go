package account

import (
	"database/sql"
	"log/slog"
)

type Repository interface {
	GetAllAccounts() ([]Account, error)
}

type DefaultRepository struct {
	Logger *slog.Logger
	Db     *sql.DB
}

// IndicateHealth to make repository a commons.HealthIndicator
func (repo *DefaultRepository) IndicateHealth() (string, bool) {
	name := "AccountRepository"
	err := repo.Db.Ping()
	if err != nil {
		repo.Logger.Error("Error connecting to database", slog.String("error", err.Error()))
		return name, false
	}
	return name, true
}

func (repo *DefaultRepository) GetAllAccounts() ([]Account, error) {

	// TODO: Replace with real implementation, when Flyway (or alt) configured
	return []Account{}, nil
}
