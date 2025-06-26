package account

import (
	"errors"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

// ===== Mocking
type MockRepository struct {
	accounts []Account
	err      error
}

func (m MockRepository) GetAllAccounts() ([]Account, error) {
	return m.accounts, m.err
}

// ===== Tests
func TestService_GetAllAccounts(t *testing.T) {
	// Given
	repo := MockRepository{
		accounts: []Account{
			{
				Id:    uuid.New(),
				Alias: "Mock Account",
			},
		},
		err: nil,
	}
	service := DefaultService{repo}

	// When
	accounts, err := service.GetAllAccounts()

	// Then
	if err != nil {
		t.Errorf("Error returned unexpectedly %s", err)
	}
	if !reflect.DeepEqual(accounts, repo.accounts) {
		t.Errorf("Accounts don't match!")
	}
}

func TestService_GetAllAccountsError(t *testing.T) {
	// Given
	repo := MockRepository{
		accounts: nil,
		err:      errors.New("some repo err"),
	}
	service := DefaultService{repo}

	// When
	_, err := service.GetAllAccounts()

	// Then
	if err == nil {
		t.Errorf("No err returned when err expected")
	}
}
