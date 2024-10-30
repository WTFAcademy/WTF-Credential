package model

import "time"

type Course struct {
	Id          int64      `gorm:"column:id;type:bigint;primaryKey;autoIncrement" json:"id"`                                    // 课程的唯一标识符
	Sort        int64      `gorm:"column:sort;type:bigint;not null" json:"sort"`                                                // 课程的排序值
	Path        string     `gorm:"column:path;type:text;not null;default:''" json:"path"`                                       // 课程的路径（通常是文件系统路径或URL）
	Repo        string     `gorm:"column:repo;type:text;not null;default:''" json:"repo"`                                       // 课程的代码仓库（如GitHub链接）
	Title       string     `gorm:"column:title;type:text;not null;unique;default:''" json:"title"`                              // 课程的标题
	Cover       string     `gorm:"column:cover;type:text;not null;default:''" json:"cover"`                                     // 课程封面图像的URL
	Level       string     `gorm:"column:level;type:text;not null;default:''" json:"level"`                                     // 课程的难度级别（如初级、中级、高级）
	Outline     string     `gorm:"column:outline;type:text;not null;default:''" json:"outline"`                                 // 课程大纲
	Category    string     `gorm:"column:category;type:text;not null;default:''" json:"category"`                               // 课程所属类别
	Portrait    string     `gorm:"column:portrait;type:text;not null;default:''" json:"portrait"`                               // 课程讲师的头像图像URL
	Schedule    string     `gorm:"column:schedule;type:text;not null;default:''" json:"schedule"`                               // 课程的安排或时间表
	Description string     `gorm:"column:description;type:text;not null;default:''" json:"description"`                         // 课程的详细描述
	CreatedAt   *time.Time `gorm:"column:created_at;type:timestamp without time zone;not null;default:now()" json:"created_at"` // 课程创建时间
	UpdatedAt   *time.Time `gorm:"column:updated_at;type:timestamp without time zone;not null;default:now()" json:"updated_at"` // 课程最后更新时间
	UserID      int64      `gorm:"column:user_id;type:bigint;not null;index:idx_courses_user_id" json:"user_id"`                // 创建该课程的用户ID
}

// TableName 指定表名为 courses
func (Course) TableName() string {
	return "courses"
}
