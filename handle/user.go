package handle

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"wtf-credential/errors"
	"wtf-credential/middleware"
	"wtf-credential/request"
	"wtf-credential/response"
	"wtf-credential/service"
)

func SignatureVerify(from, sigHex string, msg []byte) bool {
	sig := hexutil.MustDecode(sigHex)
	msg = accounts.TextHash(msg)
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}
	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}
	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	fmt.Printf("recoveredAddr %#v", recoveredAddr)
	return from == recoveredAddr.Hex()
}

func GetUserWallet(ctx *gin.Context) {
	loginUid, ok := middleware.GetUuidFromContext(ctx)
	if !ok {
		ctx.JSON(200, errors.Entity("failed to retrieve user from context"))
		return
	}
	data, err := service.GetUserWallet(ctx, loginUid)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}

func BindWallet(ctx *gin.Context) {
	loginUid, ok := middleware.GetUuidFromContext(ctx)
	if !ok {
		ctx.JSON(200, errors.Entity("failed to retrieve user from context"))
		return
	}
	var bindWalletRequest request.BindWalletRequest
	if err := ctx.ShouldBindJSON(&bindWalletRequest); err != nil {
		ctx.JSON(200, errors.Entity("param error"))
		return
	}
	if verify := SignatureVerify(bindWalletRequest.Wallet, bindWalletRequest.SignData, []byte(bindWalletRequest.MesData)); !verify {
		ctx.JSON(200, errors.Entity("signature verify error"))
		return
	}
	data, err := service.BindWallet(ctx, bindWalletRequest, loginUid)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}

func ChangeWallet(ctx *gin.Context) {
	loginUid, ok := middleware.GetUuidFromContext(ctx)
	if !ok {
		ctx.JSON(200, errors.Entity("failed to retrieve user from context"))
		return
	}
	var changeWalletRequest request.ChangeWalletRequest
	if err := ctx.ShouldBindJSON(&changeWalletRequest); err != nil {
		ctx.JSON(200, errors.Entity("param error"))
		return
	}
	if verify := SignatureVerify(changeWalletRequest.Wallet, changeWalletRequest.SignData, []byte(changeWalletRequest.MesData)); !verify {
		ctx.JSON(200, errors.Entity("signature verify error"))
		return
	}
	data, err := service.ChangeWallet(ctx, changeWalletRequest, loginUid)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}
