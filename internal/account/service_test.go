package account_test

import (
	"testing"

	"github.com/Harmelodic/init-microservice-go/internal/account"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Mocks

type MockRepository struct {
	accounts []account.Account
	err      error
}

func (m MockRepository) GetAllAccounts() ([]account.Account, error) {
	return m.accounts, m.err
}

func (m MockRepository) GetAccountByID(_ uuid.UUID) (*account.Account, error) {
	return &m.accounts[0], m.err
}

// Tests

func TestService_GetAllAccounts(t *testing.T) {
	t.Parallel()
	// Given
	mockRepo := MockRepository{
		accounts: []account.Account{
			{
				ID:    uuid.New(),
				Alias: "Mock Account",
			},
		},
		err: nil,
	}
	service := account.DefaultService{Repository: mockRepo}

	// When
	accounts, err := service.GetAllAccounts()

	// Then
	assert.Equal(t, mockRepo.accounts, accounts)
	assert.NoError(t, err)
}

func TestService_GetAllAccountsError(t *testing.T) {
	t.Parallel()
	// Given
	mockRepo := MockRepository{
		accounts: nil,
		err:      errMock,
	}
	service := account.DefaultService{Repository: mockRepo}

	// When
	_, err := service.GetAllAccounts()

	// Then
	assert.Error(t, err)
}
