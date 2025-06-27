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

func (repo *DefaultRepository) Name() string {
	return "AccountRepository"
}

func (repo *DefaultRepository) IsHealthy() bool {
	logger := commons.NewLogger()
	err := repo.Db.Ping()
	if err != nil {
		logger.Error("Error connecting to database", slog.String("error", err.Error()))
		return false
	}
	return true
}

func (repo *DefaultRepository) GetAllAccounts() ([]Account, error) {

	// TODO: Replace with real implementation
	return []Account{}, nil
}
