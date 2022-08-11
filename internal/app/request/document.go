package request

type SaveSignaturalRequest struct {
	Phone      string `json:"phone"`
	Signatural []byte `json:"signatual"`
}

type GetSignRequest struct {
	Phone  string `json:"phone"`
	Number int    `json:"number"`
}
