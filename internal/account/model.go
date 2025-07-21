package account

import "github.com/google/uuid"

// Account represents a simple account.
type Account struct {
	ID    uuid.UUID `db:"id"    json:"id"`
	Alias string    `db:"alias" json:"alias"`
}
