package daos

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/spruceid/siwe-go"
	"gorm.io/gorm"
	"time"
	model "wtf-credential/models"
)

func GenerateNonce(ctx context.Context, wallet string) (res string, err error) {
	var nonce model.TbNonce
	if err := DB.WithContext(ctx).Where("address = ?", wallet).First(&nonce).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return ``, err
		}
	}

	if &nonce != nil {
		nonce.Nonce = siwe.GenerateNonce()
		nonce.ExpiresAt = time.Now().Add(time.Minute * 5)
		err = UpdateNonce(ctx, &nonce)
	} else {
		nonce = model.TbNonce{
			Id:        uuid.New(),
			Address:   wallet,
			Nonce:     siwe.GenerateNonce(),
			Provider:  `eth`,
			ExpiresAt: time.Now().Add(time.Minute * 5), // 暂且不用
		}
		err = CreateNonce(ctx, &nonce)
	}

	if err != nil {
		return ``, err
	}
	return nonce.Nonce, err

}

// VerifyNonce 验证给定钱包地址的 nonce 是否有效
func VerifyNonce(ctx context.Context, wallet string, nonce string) (bool, error) {
	var storedNonce model.TbNonce

	// 查询数据库中该用户的 nonce
	err := DB.WithContext(ctx).Where("address = ? AND nonce = ?", wallet, nonce).First(&storedNonce).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil // 如果未找到记录，返回 false 表示 nonce 无效
		}
		return false, fmt.Errorf("failed to verify nonce: %w", err) // 其他错误
	}

	// 检查 nonce 是否过期
	if time.Now().After(storedNonce.ExpiresAt) {
		return false, nil // 如果 nonce 已过期，返回 false
	}

	return true, nil // nonce 有效
}

func CreateNonce(ctx context.Context, request *model.TbNonce) error {
	return DB.WithContext(ctx).Create(request).Error
}

func UpdateNonce(ctx context.Context, request *model.TbNonce) error {
	return DB.WithContext(ctx).Save(request).Error
}
