package account_test

import (
	"log/slog"
	"strconv"
	"testing"

	"github.com/Harmelodic/init-microservice-go/internal/account"
	"github.com/Harmelodic/init-microservice-go/internal/commons"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDefaultRepository_GetAllAccountsEmpty(t *testing.T) {
	t.Parallel()
	// Given
	logger := slog.New(slog.DiscardHandler)

	database, cleanUp := commons.NewMockDb(t, "../../migrations", logger)
	defer cleanUp()

	repository := account.DefaultRepository{
		Db: database,
	}

	// When
	resultingAccounts, err := repository.GetAllAccounts()

	// Then
	assert.Equal(t, make([]account.Account, 0), resultingAccounts) // Ensure slice is empty, but not `nil` empty
	assert.NoError(t, err)
}

func TestDefaultRepository_GetAllAccounts(t *testing.T) {
	t.Parallel()
	// Given
	logger := slog.New(slog.DiscardHandler)

	database, cleanUp := commons.NewMockDb(t, "../../migrations", logger)
	defer cleanUp()

	repository := account.DefaultRepository{
		Db: database,
	}

	accounts := make([]account.Account, 10)
	for i := range 10 {
		accounts[i] = account.Account{
			ID:    uuid.New(),
			Alias: "Account " + strconv.Itoa(i),
		}
	}

	for i := range 10 {
		_, err := database.NamedExec("INSERT INTO account VALUES (:id, :alias)", accounts[i])
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	// When
	resultingAccounts, err := repository.GetAllAccounts()

	// Then
	assert.Equal(t, accounts, resultingAccounts)
	assert.NoError(t, err)
}

func TestDefaultRepository_GetAllAccountsError(t *testing.T) {
	t.Parallel()

	logger := slog.New(slog.DiscardHandler)
	database, cleanUp := commons.NewMockDb(t, "../../migrations", logger)
	repository := account.DefaultRepository{
		Db: database,
	}

	cleanUp() // Clean up database before using it to induce connection error

	accounts, err := repository.GetAllAccounts()

	assert.Nil(t, accounts)
	assert.Error(t, err)
}

func TestDefaultRepository_GetAccountById(t *testing.T) {
	t.Parallel()
	// Given
	logger := slog.New(slog.DiscardHandler)

	database, cleanUp := commons.NewMockDb(t, "../../migrations", logger)
	defer cleanUp()

	accounts := make([]account.Account, 10)
	for i := range 10 {
		accounts[i] = account.Account{
			ID:    uuid.New(),
			Alias: "Account " + strconv.Itoa(i),
		}
	}

	for i := range 10 {
		_, err := database.NamedExec("INSERT INTO account VALUES (:id, :alias)", accounts[i])
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	repository := account.DefaultRepository{
		Db: database,
	}

	// When
	returnedAccount, err := repository.GetAccountByID(accounts[3].ID)

	// Then
	assert.Equal(t, &accounts[3], returnedAccount)
	assert.NoError(t, err)
}

func TestDefaultRepository_GetAccountByIdError(t *testing.T) {
	t.Parallel()
	// Given
	logger := slog.New(slog.DiscardHandler)
	database, cleanUp := commons.NewMockDb(t, "../../migrations", logger)
	repository := account.DefaultRepository{
		Db: database,
	}

	cleanUp() // Clean up database before using it to induce connection error

	// When
	accounts, err := repository.GetAccountByID(uuid.New())

	// Then
	assert.Nil(t, accounts)
	assert.Error(t, err)
}
