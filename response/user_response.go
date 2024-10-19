package response

type GetUserWallet struct {
	Wallet string `json:"wallet"`
}

type BindWallet struct {
	Wallet string `json:"wallet"`
}

type ChangeWallet struct {
	Wallet string `json:"wallet"`
}
