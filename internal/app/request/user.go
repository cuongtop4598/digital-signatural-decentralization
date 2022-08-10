package request

// Todo: add middleware to validate request model or add validate tag to struct => google search
type UserInfo struct {
	Name        string `json:"name"`
	Gmail       string `json:"gmail"`
	CardID      string `json:"card_id"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"date_of_birth"`
	PublicKey   string `json:"public_key"`
}

type Login struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (u *UserInfo) SantisizePassword() {
	u.Password = ""
}
