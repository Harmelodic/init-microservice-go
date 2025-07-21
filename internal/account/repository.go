package account

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Repository represents what is needed for an account repository.
type Repository interface {
	GetAllAccounts() ([]Account, error)
	GetAccountByID(uuid uuid.UUID) (*Account, error)
}

// DefaultRepository is the default SQL implementation of a Repository.
type DefaultRepository struct {
	Db *sqlx.DB
}

// GetAllAccounts fetches all accounts in the repository.
func (repo *DefaultRepository) GetAllAccounts() ([]Account, error) {
	accounts := make([]Account, 0)

	err := repo.Db.Select(&accounts, "SELECT id, alias FROM account")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all accounts from repository: %w", err)
	}

	return accounts, nil
}

// GetAccountByID fetches the account with the given ID.
func (repo *DefaultRepository) GetAccountByID(id uuid.UUID) (*Account, error) {
	account := Account{}

	err := repo.Db.Get(&account, "SELECT id, alias FROM account WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch account %s from repository: %w", id.String(), err)
	}

	return &account, nil
}
