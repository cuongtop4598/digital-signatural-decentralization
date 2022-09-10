package response

type DocumentResponse struct {
	IndexOnchain    int    `json:"index_onchain"`
	Owner           string `json:"owner"`
	Name            string `json:"name"`
	BlockNumber     string `json:"block_number,omitempty"`
	BlockHash       string `json:"block_hash,omitempty"`
	TransactionHash string `json:"transaction_hash"`
	Signature       string `json:"signature"`
	Public          bool   `json:"public"`
	TypeFile        string `json:"type_file"`
}

type ListDocumentResponse struct {
	Documents []DocumentResponse `json:"documents"`
}
