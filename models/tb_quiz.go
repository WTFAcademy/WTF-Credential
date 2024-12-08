package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// TbQuiz 表示测验的基本信息
type TbQuiz struct {
	Id         int       `json:"id"`          // 测验的唯一标识符
	Uid        uuid.UUID `json:"uid"`         // 创建该测验的用户的唯一标识符
	TotalScore int       `json:"total_score"` // 测验的总分数
	Status     int       `json:"status"`      // 测验的状态（例如：进行中、已完成、已取消等）
	CreatedAt  time.Time `json:"created_at"`  // 测验记录的创建时间
	UpdatedAt  time.Time `json:"updated_at"`  // 测验记录的最后更新时间
}

func (TbQuiz) TableName() string {
	return "tb_quiz"
}

type TbQuizExercise struct {
	Id           int       `json:"id"`
	QuizId       int       `json:"quiz_id"`
	ExerciseId   int       `json:"exercise_id"`
	Uid          uuid.UUID `json:"uid"`
	MetaData     Exercise  `gorm:"type:jsonb" json:"meta_data"`     // 元数据
	QuestionData jsonb     `gorm:"type:jsonb" json:"question_data"` // 题目数据
	Status       int       `json:"status"`                          // 状态
	CreatedAt    time.Time `json:"created_at"`                      // 创建时间
	UpdatedAt    time.Time `json:"updated_at"`                      // 更新时间
}

type TbQuizExerciseOriginal struct {
	Id           int       `json:"id"`
	QuizId       int       `json:"quiz_id"`
	ExerciseId   int       `json:"exercise_id"`
	Uid          uuid.UUID `json:"uid"`
	MetaData     jsonb     `gorm:"type:jsonb" json:"meta_data"`     // 元数据
	QuestionData jsonb     `gorm:"type:jsonb" json:"question_data"` // 题目数据
	Status       int       `json:"status"`                          // 状态
	CreatedAt    time.Time `json:"created_at"`                      // 创建时间
	UpdatedAt    time.Time `json:"updated_at"`                      // 更新时间
}

type jsonb map[string]interface{}

func (a jsonb) Value() ([]byte, error) {
	return json.Marshal(a)
}
func (a *jsonb) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), a)
}

func (TbQuizExercise) TableName() string {
	return "tb_quiz_exercise"
}
func (TbQuizExerciseOriginal) TableName() string {
	return "tb_quiz_exercise"
}

type TbLessonQuiz struct {
	Id         int       `json:"id"`
	QuizId     int       `json:"quiz_id"`
	CourseId   uuid.UUID `json:"course_id"`   // 课程ID
	LessionId  uuid.UUID `json:"lession_id"`  // 单元ID
	IsSelected int       `json:"is_selected"` // 是否被选中
	Status     int       `json:"status"`      // 状态
	CreatedAt  time.Time `json:"created_at"`  // 创建时间
	UpdatedAt  time.Time `json:"updated_at"`  // 更新时间
	Uid        uuid.UUID `json:"uid"`         // 用户ID
}

func (TbLessonQuiz) TableName() string {
	return "tb_lesson_quiz"
}

type TbLessonSelectedQuiz struct {
	Id        int       `json:"id"`
	QuizId    int       `json:"quiz_id"`
	CourseId  uuid.UUID `json:"course_id"`  // 课程ID
	LessionId uuid.UUID `json:"lession_id"` // 单元ID
	Status    int       `json:"status"`     // 状态
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}

func (TbLessonSelectedQuiz) TableName() string {
	return "tb_lesson_selected_quiz"
}

// 标记用户权限
type TbQuizRole struct {
	Id       int       `json:"id"`
	Uid      uuid.UUID `json:"uid"`       // 用户ID
	CourseId uuid.UUID `json:"course_id"` // 课程ID
	Role     int       `json:"role"`      // 角色
}

func (TbQuizRole) TableName() string {
	return "tb_quiz_role"
}

// quiz初始化
type TbQuizInit struct {
	Exercises []Exercise `json:"exercises"` // 练习题
	LessonId  uuid.UUID  `json:"lesson_id"` // 单元ID
	QuizId    int        `json:"quiz_id"`   // 测验ID
}

// editor返回结构
type TbQuizResponse struct {
	Exercises []ExerciseNoId `json:"exercises"` // 练习题
}
type Exercise struct {
	Title   string  `json:"title"`   // 标题
	Meta    Meta    `json:"meta"`    // 元数据
	Content Content `json:"content"` // 内容
}
type ExerciseNoId struct {
	Title   string   `json:"title"`   // 标题
	Meta    MetaNoId `json:"meta"`    // 元数据
	Content Content  `json:"content"` // 内容
}

type Meta struct {
	Type   string   `json:"type"`   // 类型
	Answer []string `json:"answer"` // 答案
	Score  int      `json:"score"`  // 分数
	ID     int      `json:"id"`     // ID
}
type MetaNoId struct {
	Type   string   `json:"type"`   // 类型
	Answer []string `json:"answer"` // 答案
	Score  int      `json:"score"`  // 分数
}
type Content struct {
	Extends []Extend `json:"extend"`  // 扩展
	Options []Option `json:"options"` // 选项
}
type Option struct {
	Label string `json:"label"` // 标签
	Value string `json:"value"` // 值
}
type Extend struct {
	Type string `json:"type"` // 类型
	Raw  string `json:"raw"`  // 原始数据
}

// reviewer返回结构
type TbQuizReviewResponse struct {
	Exercises []Exercise             `json:"exercises"` // 练习题
	User      UserQuizReviewResponse `json:"user"`      // 用户信息
}
type UserQuizReviewResponse struct {
	Uid      uuid.UUID `json:"uid"`       // 用户ID
	Username string    `json:"user_name"` // 用户名
}
