package account

import (
	"github.com/Harmelodic/init-microservice-go/internal/commons"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
)

// TODO when Flyway (or alt) configured and SQL table(s) created
func TestDefaultRepository_GetAllAccounts(t *testing.T) {
	logger := slog.New(slog.DiscardHandler)
	database, cleanUp := commons.NewMockDb(t, "postgres")
	defer cleanUp()
	repository := DefaultRepository{
		Logger: logger,
		Db:     database,
	}
	// TODO: Insert some entries into the DB.

	accounts, err := repository.GetAllAccounts()
	if err != nil {
		t.Error("Failed to get all accounts")
	}

	assert.NotNil(t, accounts)
	// TODO: Assert entries are the same
}

func TestDefaultRepository_GetAllAccountsError(t *testing.T) {
	// TODO: Unskip
	t.Skip("Skipping for now whilst still work in progress")
	logger := slog.New(slog.DiscardHandler)
	database, cleanUp := commons.NewMockDb(t, "postgres")
	repository := DefaultRepository{
		Logger: logger,
		Db:     database,
	}
	cleanUp() // Clean up database before using it to induce error

	accounts, err := repository.GetAllAccounts()

	assert.Nil(t, accounts)
	assert.NotNil(t, err)
}
