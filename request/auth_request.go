package request

type NonceRequest struct {
	Wallet string `json:"wallet"`
}

type GithubLoginRequest struct {
	Code string `json:"code"`
}
type LoginRequest struct {
	Message   Message `json:"message"`
	Signature string  `json:"signature"`
}
type Message struct {
	Domain         string `json:"domain"`
	Address        string `json:"address"`
	Uri            string `json:"uri"`
	Version        string `json:"version"`
	Statement      string `json:"statement"`
	Nonce          string `json:"nonce"`
	ChainID        int    `json:"chainId"`
	IssuedAt       string `json:"issuedAt"`
	ExpirationTime string `json:"expirationTime"`
}
