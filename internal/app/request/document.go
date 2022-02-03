package request

type SignDocumentRequest struct {
	UserID    string `json:"user_id"`
	DocData   []byte `json:"doc_data"`
	Signature []byte `json:"signature"`
}
