package model

import (
	"github.com/google/uuid"
)

type TbUserCore struct {
	Id       uuid.UUID `json:"id,omitempty"`     // 用户的唯一标识符 (UUID)，如果为空则省略
	Email    string    `json:"email,omitempty"`  // 用户的电子邮件地址，可能为空
	Github   string    `json:"github,omitempty"` // 用户的 GitHub 账号链接，可能为空
	UserName string    `json:"username"`         // 用户名，必填项
	Twitter  string    `json:"twitter"`          // 用户的 Twitter 账号链接
	Bio      string    `json:"bio"`              // 用户的个人简介
	Avatar   string    `json:"avatar"`           // 用户头像的链接地址
	Viewer   string    `json:"viewer"`           // 用于存储用户的查看者信息，可能是查看者 ID 或查看者角色
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
