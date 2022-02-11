package request

type UserInfo struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	CardID      string `json:"card_id"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"date_of_birth"`
	PublicKey   string `json:"public_key"`
}
