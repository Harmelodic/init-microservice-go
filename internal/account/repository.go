package account

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type Repository interface {
	GetAllAccounts() ([]Account, error)
	GetAccountById(uuid uuid.UUID) (*Account, error)
}

type DefaultRepository struct {
	Logger *slog.Logger
	Db     *sqlx.DB
}

func (repo *DefaultRepository) GetAllAccounts() ([]Account, error) {
	accounts := make([]Account, 0)
	err := repo.Db.Select(&accounts, "SELECT id, alias FROM account")
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (repo *DefaultRepository) GetAccountById(id uuid.UUID) (*Account, error) {
	account := Account{}
	err := repo.Db.Get(&account, "SELECT id, alias FROM account WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
