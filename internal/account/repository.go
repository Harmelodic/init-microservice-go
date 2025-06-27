package account

import (
	"database/sql"
	"github.com/Harmelodic/init-microservice-go/internal/commons"
	"log/slog"
)

type Repository interface {
	GetAllAccounts() ([]Account, error)
}

type DefaultRepository struct {
	Db *sql.DB
}

// IndicateHealth to make repository a commons.HealthIndicator
func (repo *DefaultRepository) IndicateHealth() (string, bool) {
	name := "AccountRepository"
	err := repo.Db.Ping()
	if err != nil {
		logger := commons.NewLogger()
		logger.Error("Error connecting to database", slog.String("error", err.Error()))
		return name, false
	}
	return name, true
}

func (repo *DefaultRepository) GetAllAccounts() ([]Account, error) {

	// TODO: Replace with real implementation
	return []Account{}, nil
}
