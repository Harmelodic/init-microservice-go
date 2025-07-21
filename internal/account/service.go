package account

import "fmt"

// Service contains the domain logic for the account package.
type Service interface {
	GetAllAccounts() ([]Account, error)
}

// DefaultService is the default implementation of Service.
type DefaultService struct {
	Repository Repository
}

// GetAllAccounts fetches all the accounts from wherever they are stored.
func (service *DefaultService) GetAllAccounts() ([]Account, error) {
	accounts, err := service.Repository.GetAllAccounts()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all accounts: %w", err)
	}

	return accounts, nil
}
