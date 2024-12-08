package model

import (
	"time"

	"github.com/google/uuid"
)

// 和中台数据库一样主要存储单个的quiz
type Quiz struct {
	ID        int64     `json:"id"`         // Quiz ID
	Sort      int64     `json:"sort"`       // 排序
	Score     int64     `json:"score"`      // 分数
	Content   string    `json:"content"`    // 内容
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	UserID    int64     `json:"user_id"`    // 用户ID 是哪个后台用户给的题目
	CourseID  int64     `json:"course_id"`  // 课程ID
	ChapterID int64     `json:"chapter_id"` // 章节ID
}

// TableName 指定该结构体对应的数据库表名
func (Quiz) TableName() string {
	return "quizzes"
}

// 我会有张整的表来存储用户的初始化题目
type QuizInit struct {
	Exercises []Exercise `json:"exercises"`  // 练习题
	ChapterID int64      `json:"chapter_id"` // 单元ID
	QuizId    int64      `json:"quiz_id"`    // 测验ID
}

// TableName 指定该结构体对应的数据库表名
func (QuizInit) TableName() string {
	return "tb_quiz_init"
}

type GeneratedQuiz struct {
	Id         int64     `json:"id"`          // 测验的唯一标识符
	Uid        uuid.UUID `json:"uid"`         // 创建该测验的用户的唯一标识符  TODO:要改成不是uuid
	TotalScore int64     `json:"total_score"` // 测验的总分数
	Status     int       `json:"status"`      // 测验的状态（例如：进行中、已完成、已取消等）
	CreatedAt  time.Time `json:"created_at"`  // 测验记录的创建时间
	UpdatedAt  time.Time `json:"updated_at"`  // 测验记录的最后更新时间
}
