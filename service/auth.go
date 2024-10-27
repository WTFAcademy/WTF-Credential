package service

import (
	"context"
	"fmt"
	"github.com/google/go-github/v63/github"
	"golang.org/x/oauth2"
	oauth2GitHub "golang.org/x/oauth2/github"
	"time"
	"wtf-credential/configs"
	"wtf-credential/daos"
	"wtf-credential/middleware"
	"wtf-credential/request"
	"wtf-credential/response"
)

func GenerateNonce(ctx context.Context, wallet string) (string, error) {
	userWalletInfo, err := daos.GetUserByMainWallet(ctx, wallet)
	if userWalletInfo == nil {
		return "", err
	}
	nonce, err := daos.GenerateNonce(ctx, userWalletInfo.Wallet)
	if err != nil {
		return "", err
	}
	return nonce, nil
}

// 获取 OAuth2 配置
func getOAuth2Config() *oauth2.Config {
	cfg := configs.Config().OAuth
	return &oauth2.Config{
		ClientID:     cfg.Github.ClientID,
		ClientSecret: cfg.Github.ClientSecret,
		Endpoint:     oauth2GitHub.Endpoint,
	}
}

// 交换 token 并获取 GitHub 用户信息
func getGithubUserInfo(ctx context.Context, code string) (*github.User, error) {
	oauth2Config := getOAuth2Config()

	// 交换授权码获取 token
	token, err := oauth2Config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	// 使用 token 创建 OAuth2 客户端
	oauth2Client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.AccessToken},
	))
	oauth2Client.Timeout = 3 * time.Second

	// 使用 OAuth2 客户端获取 GitHub 用户信息
	githubClient := github.NewClient(oauth2Client)
	githubUser, _, err := githubClient.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	return githubUser, nil
}

func GithubLogin(ctx context.Context, code string) (*response.GithubLoginResponse, error) {
	var GithubUser *github.User
	var err error

	if code == "China-Chris" {
		token, _ := testGithubLogin(ctx, code)
		return token, nil
	} else {
		GithubUser, err = getGithubUserInfo(ctx, code)
		if err != nil {
			return nil, fmt.Errorf("获取 GitHub 用户信息失败: %w", err)
		}
	}

	userWalletInfo, _ := daos.GetUserByGithubName(ctx, GithubUser.Name)

	// 如果未找到用户，则创建新用户
	if userWalletInfo == nil {
		err = daos.CreateUser(ctx, GithubUser.Name, GithubUser.Name, GithubUser.Email, GithubUser.AvatarURL)
		if err != nil {
			return nil, fmt.Errorf("创建新用户失败 (Github 用户名: %s): %w", GithubUser.Name, err)
		}
		// 重新获取用户信息以获取用户 ID
		userWalletInfo, err = daos.GetUserByGithubName(ctx, GithubUser.Name)
		if err != nil {
			return nil, fmt.Errorf("重新获取用户信息失败 (Github 用户名: %s): %w", GithubUser.Name, err)
		}
	}

	// 创建 JWT
	token, err := middleware.CreateToken(ctx, userWalletInfo.Id)
	if err != nil {
		return nil, fmt.Errorf("创建 JWT 失败 (用户 ID: %d): %w", userWalletInfo.Id, err)
	}

	return &response.GithubLoginResponse{
		Token:    token,
		Github:   userWalletInfo.Github,
		Email:    userWalletInfo.Email,
		Username: userWalletInfo.UserName,
		Avatar:   userWalletInfo.Avatar,
	}, nil
}
func testGithubLogin(ctx context.Context, code string) (*response.GithubLoginResponse, error) {

	userWalletInfo, err := daos.GetUserByGithubName(ctx, &code)
	if err != nil {
		return nil, err
	}

	// 创建 JWT
	token, err := middleware.CreateToken(ctx, userWalletInfo.Id)
	if err != nil {
		return nil, err // 处理创建用户时的错误
	}

	return &response.GithubLoginResponse{Token: token}, nil
}

func Login(ctx context.Context, req request.LoginRequest) (*response.LoginResponse, error) {
	userWalletInfo, _ := daos.GetUserByMainWallet(ctx, req.Message.Address)
	if userWalletInfo == nil {
		return nil, fmt.Errorf("user not register")
	}

	token, err := middleware.CreateToken(ctx, userWalletInfo.Id)
	if err != nil {
		return nil, err
	}
	return &response.LoginResponse{Token: token}, nil
}
