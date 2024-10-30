package model

import (
	"github.com/google/uuid"
	"time"
)

// TbLessonQuizGrade 表示用户在课程测验中的成绩信息
type TbLessonQuizGrade struct {
	Id        int64     `json:"id,omitempty"`         // 成绩记录的唯一标识符
	Uid       uuid.UUID `json:"uid,omitempty"`        // 用户的唯一标识符，指明该成绩属于哪个用户
	QuizId    string    `json:"quiz_id,omitempty"`    // 测验的唯一标识符，指明该成绩对应的测验
	Email     string    `json:"email,omitempty"`      // 用户的电子邮件地址，用于联系或识别用户
	Score     int64     `json:"score,omitempty"`      // 用户在测验中的得分
	LessonId  uuid.UUID `json:"lesson_id,omitempty"`  // 关联的课程单元ID，指明该成绩属于哪个课程单元
	CourseId  uuid.UUID `json:"course_id,omitempty"`  // 关联的课程ID，指明该成绩属于哪个课程
	Progress  int64     `json:"progress"`             // 用户在测验中的完成进度（百分比）
	CreatedAt time.Time `json:"created_at,omitempty"` // 成绩记录的创建时间
	UpdatedAt time.Time `json:"updated_at,omitempty"` // 成绩记录的最后更新时间
}

func (TbLessonQuizGrade) TableName() string {
	return "tb_lesson_quiz_grade"
}
