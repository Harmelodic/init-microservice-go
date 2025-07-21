package account_test

import (
	"encoding/json"
	"github.com/Harmelodic/init-microservice-go/internal/account"
	"github.com/Harmelodic/init-microservice-go/internal/commons"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mocks

type MockService struct {
	accounts []account.Account
	err      error
}

func (m MockService) GetAllAccounts() ([]account.Account, error) {
	return m.accounts, m.err
}

// Tests

func TestController_GetAllAccounts(t *testing.T) {
	t.Parallel()
	// Given
	testEngine := commons.NewGinEngine("test", slog.New(slog.DiscardHandler))
	mockService := MockService{
		accounts: []account.Account{
			{
				ID:    uuid.New(),
				Alias: "Mock Account",
			},
		},
		err: nil,
	}
	account.Controller(testEngine, mockService, slog.New(slog.DiscardHandler))

	accountJSON, err := json.Marshal(mockService.accounts)
	if err != nil {
		t.Fatalf("Could not marshall JSON: %s", err.Error())
	}

	// When
	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/account", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	// Then
	assert.Equal(t, 200, responseRecorder.Code)
	assert.JSONEq(t, string(accountJSON), responseRecorder.Body.String())
}

func TestController_GetAllAccountsError(t *testing.T) {
	t.Parallel()
	// Given
	testEngine := commons.NewGinEngine("test", slog.New(slog.DiscardHandler))
	mockService := MockService{
		accounts: nil,
		err:      errMock,
	}
	account.Controller(testEngine, mockService, slog.New(slog.DiscardHandler))

	// When
	responseRecorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/account", http.NoBody)
	testEngine.ServeHTTP(responseRecorder, req)

	// Then
	assert.Equal(t, 500, responseRecorder.Code)
	assert.Empty(t, responseRecorder.Body.String()) // Empty body on error for security
}
