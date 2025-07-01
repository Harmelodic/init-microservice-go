package account

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Mocks

type MockRepository struct {
	accounts []Account
	err      error
}

func (m MockRepository) GetAllAccounts() ([]Account, error) {
	return m.accounts, m.err
}

func (m MockRepository) GetAccountById(_ uuid.UUID) (*Account, error) {
	return &m.accounts[0], m.err
}

// Tests

func TestService_GetAllAccounts(t *testing.T) {
	// Given
	mockRepo := MockRepository{
		accounts: []Account{
			{
				Id:    uuid.New(),
				Alias: "Mock Account",
			},
		},
		err: nil,
	}
	service := DefaultService{mockRepo}

	// When
	accounts, err := service.GetAllAccounts()

	// Then
	assert.Equal(t, mockRepo.accounts, accounts)
	assert.NoError(t, err)
}

func TestService_GetAllAccountsError(t *testing.T) {
	// Given
	mockRepo := MockRepository{
		accounts: nil,
		err:      errors.New("some mockRepo err"),
	}
	service := DefaultService{mockRepo}

	// When
	_, err := service.GetAllAccounts()

	// Then
	assert.Error(t, err)
}
