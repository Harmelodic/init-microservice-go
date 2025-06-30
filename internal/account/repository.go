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
	// TODO: Replace with real implementation, when DB migrations configured
	return []Account{}, nil
}
