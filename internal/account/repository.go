package account

import (
	"github.com/google/uuid"
)

type Repository interface {
	GetAllAccounts() ([]Account, error)
}

type DefaultRepository struct{}

func (*DefaultRepository) GetAllAccounts() ([]Account, error) {
	// TODO: Replace with DB connection
	accounts := []Account{
		{
			Id:    uuid.New(),
			Alias: "Account 1",
		},
		{
			Id:    uuid.New(),
			Alias: "Account 2",
		},
		{
			Id:    uuid.New(),
			Alias: "Account 3",
		}}
	return accounts, nil
}

func (repo *DefaultRepository) Name() string {
	return "AccountRepository"
}

func (repo *DefaultRepository) IsHealthy() bool {
	return true // TODO: Change to check DB connection health
}
