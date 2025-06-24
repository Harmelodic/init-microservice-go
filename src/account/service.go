package account

import "github.com/google/uuid"

type Service struct{}

func (Service) GetAllAccounts() []Account {
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
