package response

type GetUserWallet struct {
	Wallet string `json:"wallet"` // 用户的钱包地址
}

type BindWallet struct {
	Wallet string `json:"wallet"` // 绑定到用户账户的钱包地址
}

type ChangeWallet struct {
	Wallet string `json:"wallet"` // 更新到用户档案的新钱包地址
}

type GetProfileByUserID struct {
	Github   string `json:"github"`   // 用户的 GitHub 用户名
	Email    string `json:"email"`    // 用户的电子邮件地址
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 用户昵称
	Twitter  string `json:"twitter"`  // 用户的 Twitter 账号
	Bio      string `json:"bio"`      // 用户简介或描述
	Viewer   string `json:"viewer"`   // 查看者与用户的关系状态
	Avatar   string `json:"avatar"`   // 用户头像图片的 URL
	Wallet   string `json:"wallet"`   // 与用户账户关联的钱包地址
}
