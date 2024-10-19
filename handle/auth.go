package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/spruceid/siwe-go"
	"strconv"
	"wtf-credential/daos"
	"wtf-credential/errors"
	"wtf-credential/request"
	"wtf-credential/response"
	"wtf-credential/service"
)

func GithubLogin(ctx *gin.Context) {
	var githubLoginRequest request.GithubLoginRequest
	if err := ctx.ShouldBindJSON(&githubLoginRequest); err != nil {
		ctx.JSON(200, errors.Entity("param error"))
		return
	}
	data, err := service.GithubLogin(ctx, githubLoginRequest.Code)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}

func handleNonceValidation(ctx *gin.Context, wallet string, nonce string) bool {
	valid, err := daos.VerifyNonce(ctx, wallet, nonce)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return false
	}
	if !valid {
		ctx.JSON(200, errors.Entity("invalid nonce or nonce expired"))
		return false
	}
	return true
}

func handleSignatureVerification(ctx *gin.Context, message request.Message, signature string) bool {
	safe, err := SiweSignatureVerify(message, signature)
	if err != nil || !safe {
		ctx.JSON(200, errors.Entity("signature error"))
		return false
	}
	return true
}
func SiweSignatureVerify(message request.Message, signature string) (bool, error) {
	// jwt 登录流程
	siweMessage, err := siwe.InitMessage(message.Domain, message.Address, message.Uri, message.Nonce, map[string]interface{}{
		"statement":      message.Statement,
		"issuedAt":       message.IssuedAt,
		"nonce":          message.Nonce,
		"chainId":        strconv.Itoa(message.ChainID),
		"expirationTime": message.ExpirationTime,
	})
	if err != nil {
		return false, err
	}
	_, err = siweMessage.Verify(signature, nil, &message.Nonce, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

func Login(ctx *gin.Context) {
	var loginRequest request.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(200, errors.Entity("param error"))
		return
	}
	if !handleNonceValidation(ctx, loginRequest.Message.Address, loginRequest.Message.Nonce) {
		return
	}
	if !handleSignatureVerification(ctx, loginRequest.Message, loginRequest.Signature) {
		return
	}
	data, err := service.Login(ctx, loginRequest)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}

// 生成Nonce
func GenerateNonce(ctx *gin.Context) {
	var nonceRequest request.NonceRequest
	if err := ctx.ShouldBindJSON(&nonceRequest); err != nil {
		ctx.JSON(200, errors.Entity("param error"))
		return
	}
	data, err := service.GenerateNonce(ctx, nonceRequest.Wallet) //调用 server层
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}
