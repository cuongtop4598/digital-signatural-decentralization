package request

// Todo: add middleware to validate request model or add validate tag to struct => google search
type UserInfo struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	CardID      string `json:"card_id"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"date_of_birth"`
	PublicKey   string `json:"public_key"`
}

type Login struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
