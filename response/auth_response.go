package response

type GithubLoginResponse struct {
	Token    string `json:"token"`
	Github   string `json:"github"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type LoginResponse struct {
	Token string `json:"token"` //token
}
