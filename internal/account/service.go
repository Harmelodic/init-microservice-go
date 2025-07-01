package account

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
		return nil, err
	}

	return accounts, nil
}
