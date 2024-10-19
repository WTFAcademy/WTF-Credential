package model

import (
	"github.com/google/uuid"
)

type TbUserCore struct {
	Id     uuid.UUID `json:"id,omitempty"`
	Email  string    `json:"email,omitempty"`
	Github string    `json:"github,omitempty"`
}

func (TbUserCore) TableName() string {
	return "tb_user_core"
}

type TbUserCoreFull struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

func (TbUserCoreFull) TableName() string {
	return "tb_user_core"
}
