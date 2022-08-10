package request

type SaveSignatural struct {
	Phone      string `json:"phone"`
	Signatural []byte `json:"signatual"`
}
