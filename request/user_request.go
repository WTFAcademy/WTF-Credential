package request

type BindWalletRequest struct {
	SignData string `json:"signData"`
	MesData  string `json:"mesData"`
	Wallet   string `json:"wallet"`
}
type ChangeWalletRequest struct {
	SignData string `json:"signData"`
	MesData  string `json:"mesData"`
	Wallet   string `json:"wallet"`
}
