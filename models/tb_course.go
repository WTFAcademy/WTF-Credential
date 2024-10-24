package model

import (
	"github.com/google/uuid"
	"time"
)

type TbCourse struct {
	Id           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	RoutePath    string    `json:"route_path"`
	Description  string    `json:"description"`
	CoverImg     string    `json:"cover_img"`
	Sort         int       `json:"sort"`
	TotalScore   int       `json:"total_score"`
	UserCnt      int       `json:"user_cnt"`
	StartStatus  int       `json:"start_status"`
	ShareUrl     string    `json:"share_url"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ExtendedInfo string    `json:"extended_info"`
}

func (TbCourse) TableName() string {
	return "tb_course"
}
