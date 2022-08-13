package response

import (
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	PublicKey   string    `json:"public_key"`
	CardID      string    `json:"card_id"`
	Phone       string    `json:"phone"`
	Gmail       string    `json:"gmail"`
	DateOfBirth string    `json:"date_of_birth"`
	Hash        string    `json:"hash_info"`
}
