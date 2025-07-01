package account

import (
	"fmt"
	"github.com/Harmelodic/init-microservice-go/internal/commons"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"strconv"
	"testing"
)

func createTable(t *testing.T, db *sqlx.DB) {
	// TODO Replace with DB migrations
	_, err := db.Exec("CREATE TABLE account (id UUID, alias VARCHAR(255))")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDefaultRepository_GetAllAccounts(t *testing.T) {
	t.Parallel()
	logger := slog.New(slog.DiscardHandler)
	database, cleanUp := commons.NewMockDb(t)
	defer cleanUp()
	createTable(t, database)
	var accounts []Account
	for i := 0; i < 10; i++ {
		accounts = append(accounts, Account{
			Id:    uuid.New(),
			Alias: fmt.Sprintf("Account %s", strconv.Itoa(i)),
		})
	}
	for i := 0; i < 10; i++ {
		_, err := database.NamedExec("INSERT INTO account VALUES (:id, :alias)", accounts[i])
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	repository := DefaultRepository{
		Logger: logger,
		Db:     database,
	}

	resultingAccounts, err := repository.GetAllAccounts()

	// Then
	assert.Equal(t, accounts, resultingAccounts)
	assert.NoError(t, err)
}

func TestDefaultRepository_GetAllAccountsError(t *testing.T) {
	t.Parallel()
	logger := slog.New(slog.DiscardHandler)
	database, cleanUp := commons.NewMockDb(t)
	repository := DefaultRepository{
		Logger: logger,
		Db:     database,
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
	database, cleanUp := commons.NewMockDb(t)
	defer cleanUp()
	createTable(t, database)
	var accounts []Account
	for i := 0; i < 10; i++ {
		accounts = append(accounts, Account{
			Id:    uuid.New(),
			Alias: fmt.Sprintf("Account %s", strconv.Itoa(i)),
		})
	}
	for i := 0; i < 10; i++ {
		_, err := database.NamedExec("INSERT INTO account VALUES (:id, :alias)", accounts[i])
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	repository := DefaultRepository{
		Logger: logger,
		Db:     database,
	}

	// When
	account, err := repository.GetAccountById(accounts[3].Id)

	// Then
	assert.Equal(t, &accounts[3], account)
	assert.NoError(t, err)
}

func TestDefaultRepository_GetAccountByIdError(t *testing.T) {
	t.Parallel()
	// Given
	logger := slog.New(slog.DiscardHandler)
	database, cleanUp := commons.NewMockDb(t)
	repository := DefaultRepository{
		Logger: logger,
		Db:     database,
	}
	cleanUp() // Clean up database before using it to induce connection error

	// When
	accounts, err := repository.GetAccountById(uuid.New())

	// Then
	assert.Nil(t, accounts)
	assert.Error(t, err)
}
