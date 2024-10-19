package daos

import (
	"context"
	"fmt"
	model "wtf-credential/models"
)

func GetWalletByUserId(ctx context.Context, wallet string) (res *model.TbWallet, err error) {
	var user = model.TbWallet{}
	if err = DB.WithContext(ctx).Where("id = ?", wallet).Where("is_active = ?", true).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, err
}

func GetUserByMainWallet(ctx context.Context, wallet string) (res *model.TbWallet, err error) {
	var user = model.TbWallet{}
	if err = DB.WithContext(ctx).Where("wallet = ?", wallet).Where("is_active = ?", true).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, err
}

func BindWallet(ctx context.Context, request model.TbWallet) error {
	if err := DB.WithContext(ctx).Create(&request).Error; err != nil {
		fmt.Printf("BindWallet error: %v\n", err)
		return err
	}
	return nil
}

func ChangeWallet(ctx context.Context, request model.TbWallet) error {
	result := DB.WithContext(ctx).Model(&request).UpdateColumn("is_active", false)
	if result.Error != nil {
		fmt.Printf("ChangeWallet error: %v\n", result.Error)
		return result.Error
	}
	return nil
}
