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
		t.Fatalf("Failed to create table: %s", err.Error())
	}
}

func TestDefaultRepository_GetAllAccounts(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	database, cleanUp := commons.NewMockDb(t)
	defer cleanUp()
	createTable(t, database)
	for i := 0; i < 10; i++ {
		_, err := database.NamedExec("INSERT INTO account VALUES (:id, :alias)", &Account{
			Id:    uuid.New(),
			Alias: fmt.Sprintf("Account %s", strconv.Itoa(i)),
		})
		if err != nil {
			t.Fatalf("Failed to insert accounts: %s", err.Error())
		}
	}
	repository := DefaultRepository{
		Logger: logger,
		Db:     database,
	}

	accounts, err := repository.GetAllAccounts()
	if err != nil {
		t.Fatalf("Failed to get all accounts: %s", err.Error())
	}

	assert.NotNil(t, accounts)
	assert.Equal(t, 10, len(accounts))
}

func TestDefaultRepository_GetAllAccountsError(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	database, cleanUp := commons.NewMockDb(t)
	repository := DefaultRepository{
		Logger: logger,
		Db:     database,
	}
	cleanUp() // Clean up database before using it to induce connection error

	accounts, err := repository.GetAllAccounts()

	assert.Nil(t, accounts)
	assert.NotNil(t, err)
	t.Logf("Error retrieved: %s", err.Error())
}
