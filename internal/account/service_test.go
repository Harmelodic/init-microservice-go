package account

import (
	"errors"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

// ===== Mocking
type MockRepository struct {
	mockAccounts []Account
	mockError    error
}

func (m MockRepository) GetAllAccounts() ([]Account, error) {
	return m.mockAccounts, m.mockError
}

// ===== Tests
func TestService_GetAllAccounts(t *testing.T) {
	// Given
	repo := MockRepository{
		mockAccounts: []Account{
			{
				Id:    uuid.New(),
				Alias: "Mock Account",
			},
		},
		mockError: nil,
	}
	service := DefaultService{repo}

	// When
	accounts, err := service.GetAllAccounts()

	// Then
	if err != nil {
		t.Errorf("Error returned unexpectedly %s", err)
	}
	if !reflect.DeepEqual(accounts, repo.mockAccounts) {
		t.Errorf("Accounts don't match!")
	}
}

func TestService_GetAllAccountsError(t *testing.T) {
	// Given
	repo := MockRepository{
		mockAccounts: nil,
		mockError:    errors.New("some repo error"),
	}
	service := DefaultService{repo}

	// When
	_, err := service.GetAllAccounts()

	// Then
	if err == nil {
		t.Errorf("No error returned when error expected")
	}
}
