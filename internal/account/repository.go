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

func (repo *DefaultRepository) GetAllAccounts() ([]Account, error) {
	// TODO: Replace with real implementation, when Flyway (or alt) configured
	return []Account{}, nil
}
