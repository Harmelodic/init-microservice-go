package account

import "errors"

// Service contains the domain logic for the account package.
type Service interface {
	GetAllAccounts() ([]Account, error)
}

type DefaultService struct {
	Repository Repository
}

func (service *DefaultService) GetAllAccounts() ([]Account, error) {
	accounts, err := service.Repository.GetAllAccounts()
	if err != nil {
		return nil, errors.New("failed to fetch accounts from repository")
	}

	return accounts, nil
}
