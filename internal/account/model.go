package account

import "github.com/google/uuid"

type Account struct {
	Id    uuid.UUID `json:"id"`
	Alias string    `json:"alias"`
}
