package account

import (
	"github.com/google/uuid"
)

type Repository struct{}

func (*Repository) GetAllAccounts() ([]Account, error) {
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

func (repo *Repository) Name() string {
	return "AccountRepository"
}

func (repo *Repository) IsHealthy() bool {
	return true // TODO: Change to check DB connection health
}
