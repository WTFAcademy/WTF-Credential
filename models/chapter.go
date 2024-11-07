package model

import "time"

// Chapter represents a chapter in a course.
type Chapter struct {
	Id         int64      `gorm:"column:id;type:bigint;primaryKey;autoIncrement" json:"id"` // 章节的唯一标识符
	Sort       int64      `gorm:"column:sort;type:bigint;not null" json:"sort"`             // 章节的排序值
	Path       string     `gorm:"column:path;type:text;not null;default:''" json:"path"`
	RoutePath  string     `gorm:"column:route_path;type:text;not null;default:''" json:"route_path"`                           // 章节的路径
	Title      string     `gorm:"column:title;type:text;not null;default:''" json:"title"`                                     // 章节的标题
	Content    string     `gorm:"column:content;type:text;not null;default:''" json:"content"`                                 // 章节的内容
	Version    int64      `gorm:"column:version;type:bigint;not null" json:"version"`                                          // 章节的版本
	Keywords   string     `gorm:"column:keywords;type:text;not null;default:''" json:"keywords"`                               // 章节的关键词
	ContentURL string     `gorm:"column:content_url;type:text;not null;default:''" json:"content_url"`                         // 章节内容的URL
	CreatedAt  *time.Time `gorm:"column:created_at;type:timestamp without time zone;not null;default:now()" json:"created_at"` // 章节创建时间
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:timestamp without time zone;not null;default:now()" json:"updated_at"` // 章节最后更新时间
	UserID     int64      `gorm:"column:user_id;type:bigint;not null;index:idx_chapter_user_id" json:"user_id"`                // 创建该章节的用户ID
	CourseID   int64      `gorm:"column:course_id;type:bigint;not null;index:idx_chapter_course_id" json:"course_id"`          // 所属课程的ID
}

// TableName specifies the table name for Chapter struct.
func (Chapter) TableName() string {
	return "chapters"
}
