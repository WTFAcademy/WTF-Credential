package model

import (
	"github.com/google/uuid"
	"time"
)

type TbWallet struct {
	Id              uuid.UUID `json:"id,omitempty"`
	TbId            uuid.UUID `json:"tb_id,omitempty"`
	Wallet          string    `json:"wallet,omitempty"`
	IsActive        bool      `json:"is_active,omitempty" gorm:"type:boolean"`
	ModifyLockUntil time.Time `json:"modify_lock_until,omitempty" gorm:"type:timestamp with time zone"` // 时间戳字段，包含时区信息
}

func (TbWallet) TableName() string {
	return "tb_wallet"
}
