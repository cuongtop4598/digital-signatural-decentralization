package request

type UserInfo struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	CardID      string `json:"idcard"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"dateofbirth"`
	PublicKey   string `json:"publickey"`
}
