package account

import (
	"github.com/Harmelodic/init-microservice-go/src/commons"
	"github.com/google/uuid"
)

// Service contains the domain logic for the account package.
type Service struct{}

func (Service) GetAllAccounts() []Account {
	logger := commons.NewLogger()

	logger.Info("No implementation yet, generating accounts in-memory")
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
	return accounts
}
