package model

import (
	"github.com/google/uuid"
	"time"
)

type JSONMap map[string]interface{}

type TbNonce struct {
	Id        uuid.UUID `json:"id,omitempty" gorm:"primaryKey"`
	Provider  string    `json:"provider,omitempty"`
	Address   string    `json:"address,omitempty"`
	Nonce     string    `json:"nonce,omitempty"`
	Message   string    `json:"message,omitempty"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (TbNonce) TableName() string {
	return "tb_nonce"
}
