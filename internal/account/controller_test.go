package account

import (
	"encoding/json"
	"errors"
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
	accounts []Account
	err      error
}

func (m MockService) GetAllAccounts() ([]Account, error) {
	return m.accounts, m.err
}

// Tests

func TestController_GetAllAccounts(t *testing.T) {
	// Given
	testEngine := commons.NewGinEngine(slog.New(slog.DiscardHandler))
	mockService := MockService{
		accounts: []Account{
			{
				Id:    uuid.New(),
				Alias: "Mock Account",
			},
		},
		err: nil,
	}
	Controller(testEngine, mockService)
	accountJson, _ := json.Marshal(mockService.accounts)

	// When
	responseRecorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/account", nil)
	testEngine.ServeHTTP(responseRecorder, req)

	// Then
	assert.Equal(t, 200, responseRecorder.Code)
	assert.Equal(t, string(accountJson), responseRecorder.Body.String())
}

func TestController_GetAllAccountsError(t *testing.T) {
	// Given
	testEngine := commons.NewGinEngine(slog.New(slog.DiscardHandler))
	mockService := MockService{
		accounts: nil,
		err:      errors.New("some service err"),
	}
	Controller(testEngine, mockService)

	// When
	responseRecorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/account", nil)
	testEngine.ServeHTTP(responseRecorder, req)

	// Then
	assert.Equal(t, 500, responseRecorder.Code)
	assert.Contains(t, responseRecorder.Body.String(), mockService.err.Error())
}
