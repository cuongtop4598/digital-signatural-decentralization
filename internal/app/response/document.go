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

type MultiSignerDocumentResponse struct {
	IndexOnchain    int    `json:"index_onchain"`
	Creater         string `json:"creater"`
	PartnerA        string `json:"partner_a"`
	PartnerB        string `json:"partner_b"`
	Name            string `json:"name"`
	BlockNumber     string `json:"block_number,omitempty"`
	BlockHash       string `json:"block_hash,omitempty"`
	TransactionHash string `json:"transaction_hash"`
	SignatureA      string `json:"signature_a"`
	SignatureB      string `json:"signature_b"`
	Digest          string `json:"digest"`
	SignDateA       string `json:"sign_date_a"`
	SignDateB       string `json:"sign_date_b"`
	Complete        bool   `json:"complete"`
	Public          bool   `json:"public"`
	TypeFile        string `json:"type_file"`
}

type ListDocumentResponse struct {
	Documents []DocumentResponse `json:"documents"`
}

type ListMultiSingerDocumentResponse struct {
	Documents []MultiSignerDocumentResponse `json:"documents"`
}
