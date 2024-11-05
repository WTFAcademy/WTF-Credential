package model

import (
	"github.com/google/uuid"
	"time"
)

type TbCourseLesson struct {
	Id            uuid.UUID `json:"id,omitempty"`             // 课程单元的唯一标识符
	Title         string    `json:"title,omitempty"`          // 课程单元的标题
	Path          string    `json:"path,omitempty"`           // 课程路由路径
	RoutePath     string    `json:"route_path,omitempty"`     // 课程单元的路由路径，用于访问该单元
	CourseId      uuid.UUID `json:"course_id,omitempty"`      // 关联的课程ID，标识该单元所属的课程
	EstimatedTime string    `json:"estimated_time,omitempty"` // 估算完成该课程单元所需的时间
	Sort          int       `json:"sort,omitempty"`           // 课程单元的排序编号，用于排列显示顺序
	CreatedAt     time.Time `json:"created_at,omitempty"`     // 课程单元的创建时间
	UpdatedAt     string    `json:"updated_at,omitempty"`     // 课程单元的最后更新时间
}

func (TbCourseLesson) TableName() string {
	return "tb_course_lesson"
}
