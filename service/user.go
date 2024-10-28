package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
	"wtf-credential/daos"
	model "wtf-credential/models"
	"wtf-credential/request"
	"wtf-credential/response"
)

func GetUserWallet(ctx context.Context, loginUid string) (*response.GetUserWallet, error) {
	userWallet, err := daos.GetWalletByUserId(ctx, loginUid)
	if err != nil {
		return nil, err
	}
	return &response.GetUserWallet{Wallet: userWallet.Wallet}, nil
}

// BindWallet
func BindWallet(ctx context.Context, req request.BindWalletRequest, loginUid string) (*response.BindWallet, error) {
	wallet := model.TbWallet{
		Id:       uuid.MustParse(loginUid),
		Wallet:   req.Wallet,
		IsActive: true,
		TbId:     uuid.New(),
	}
	err := daos.BindWallet(ctx, wallet)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// ChangeWallet
func ChangeWallet(ctx context.Context, req request.ChangeWalletRequest, loginUid string) (*response.ChangeWallet, error) {
	if err := checkWalletBinding(ctx, req.Wallet); err != nil {
		return nil, err
	}
	wallet, err := daos.GetWalletByUserId(ctx, loginUid)
	if err != nil {
		return nil, err
	}
	if err := validateWalletModification(*wallet, req.Wallet); err != nil {
		return nil, err
	}
	if err := unbindWallet(ctx, *wallet); err != nil {
		return nil, err
	}
	if err := bindNewWallet(ctx, wallet.Id, req.Wallet); err != nil {
		return nil, err
	}
	return nil, nil
}

func checkWalletBinding(ctx context.Context, wallet string) error {
	if _, err := daos.GetUserByMainWallet(ctx, wallet); err == nil {
		return fmt.Errorf("wallet has already been bound")
	}
	return nil
}

func validateWalletModification(wallet model.TbWallet, newWallet string) error {
	if !wallet.ModifyLockUntil.IsZero() && time.Now().Before(wallet.ModifyLockUntil) {
		return fmt.Errorf("抱歉，目前不能修改钱包地址，修改锁定时间为3个月")
	} else if wallet.Wallet == newWallet {
		return fmt.Errorf("与原来的钱包相同!")
	}
	return nil
}

func unbindWallet(ctx context.Context, wallet model.TbWallet) error {
	walletUnbind := model.TbWallet{
		Id:              wallet.Id,
		TbId:            wallet.TbId,
		Wallet:          wallet.Wallet,
		IsActive:        false,
		ModifyLockUntil: time.Now(),
	}
	return daos.ChangeWallet(ctx, walletUnbind)
}

func bindNewWallet(ctx context.Context, userId uuid.UUID, newWallet string) error {
	wallet := model.TbWallet{
		Id:              userId,
		TbId:            uuid.New(),
		Wallet:          newWallet,
		IsActive:        true,
		ModifyLockUntil: time.Now().AddDate(0, 3, 0),
	}
	return daos.BindWallet(ctx, wallet)
}

// 辅助函数用于判断错误是否为“未找到”错误
func isRecordNotFoundError(err error) bool {
	// 检查错误是否为“未找到记录”类型（可根据具体数据库/DAO实现更新判断逻辑）
	return err.Error() == "record not found"
}
func GetProfileByUserID(ctx context.Context, loginUid string) (*response.GetProfileByUserID, error) {
	userProfile, err := daos.GetUserProfileByID(ctx, loginUid)
	if err != nil {
		return nil, err
	}

	// 获取钱包地址，但如果没有找到钱包，不返回错误
	wallet, err := daos.GetWalletByUserId(ctx, loginUid)
	if err != nil && !isRecordNotFoundError(err) { // 判断是否为“未找到”错误
		fmt.Printf("获取钱包时出现错误: %v\n", err)
	}

	profile := &response.GetProfileByUserID{
		Github:   userProfile.Github,   // 用户的 GitHub 用户名
		Email:    userProfile.Email,    // 用户的电子邮件地址
		Username: userProfile.Username, // 用户名
		Nickname: userProfile.Nickname, // 用户昵称
		Twitter:  userProfile.Twitter,  // 用户的 Twitter 账号
		Bio:      userProfile.Bio,      // 用户简介或描述
		Viewer:   userProfile.Viewer,   // 查看者与用户的关系状态
		Avatar:   userProfile.Avatar,   // 用户头像图片的 URL
		Wallet:   "",                   // 默认值为空字符串
	}

	// 如果钱包数据存在，则更新 profile.Wallet
	if wallet != nil {
		profile.Wallet = wallet.Wallet
	}

	return profile, nil
}

// UnbindWallet
func UnbindWallet(ctx context.Context, loginUid string) error {
	wallet, err := daos.GetWalletByUserId(ctx, loginUid)
	if err != nil {
		return err
	}
	if err := unbindWallet(ctx, *wallet); err != nil {
		return err
	}
	return nil
}
