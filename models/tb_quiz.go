package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
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
	MetaData     Exercise  `gorm:"type:jsonb" json:"meta_data"`
	QuestionData jsonb     `gorm:"type:jsonb" json:"question_data"`
	Status       int       `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type TbQuizExerciseOriginal struct {
	Id           int       `json:"id"`
	QuizId       int       `json:"quiz_id"`
	ExerciseId   int       `json:"exercise_id"`
	Uid          uuid.UUID `json:"uid"`
	MetaData     jsonb     `gorm:"type:jsonb" json:"meta_data"`
	QuestionData jsonb     `gorm:"type:jsonb" json:"question_data"`
	Status       int       `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
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
	CourseId   uuid.UUID `json:"course_id"`
	LessionId  uuid.UUID `json:"lession_id"`
	IsSelected int       `json:"is_selected"`
	Status     int       `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Uid        uuid.UUID `json:"uid"`
}

func (TbLessonQuiz) TableName() string {
	return "tb_lesson_quiz"
}

type TbLessonSelectedQuiz struct {
	Id        int       `json:"id"`
	QuizId    int       `json:"quiz_id"`
	CourseId  uuid.UUID `json:"course_id"`
	LessionId uuid.UUID `json:"lession_id"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (TbLessonSelectedQuiz) TableName() string {
	return "tb_lesson_selected_quiz"
}

// 标记用户权限
type TbQuizRole struct {
	Id       int       `json:"id"`
	Uid      uuid.UUID `json:"uid"`
	CourseId uuid.UUID `json:"course_id"`
	Role     int       `json:"role"`
}

func (TbQuizRole) TableName() string {
	return "tb_quiz_role"
}

// quiz初始化
type TbQuizInit struct {
	Exercises []Exercise `json:"exercises"`
	LessonId  uuid.UUID  `json:"lesson_id"`
	QuizId    int        `json:"quiz_id"`
}

// editor返回结构
type TbQuizResponse struct {
	Exercises []ExerciseNoId `json:"exercises"`
}
type Exercise struct {
	Title   string  `json:"title"`
	Meta    Meta    `json:"meta"`
	Content Content `json:"content"`
}
type ExerciseNoId struct {
	Title   string   `json:"title"`
	Meta    MetaNoId `json:"meta"`
	Content Content  `json:"content"`
}

type Meta struct {
	Type   string   `json:"type"`
	Answer []string `json:"answer"`
	Score  int      `json:"score"`
	ID     int      `json:"id"` // 用于溯源题目
}
type MetaNoId struct {
	Type   string   `json:"type"`
	Answer []string `json:"answer"`
	Score  int      `json:"score"`
}
type Content struct {
	Extends []Extend `json:"extend"`
	Options []Option `json:"options"`
}
type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
type Extend struct {
	Type string `json:"type"`
	Raw  string `json:"raw"`
}

// reviewer返回结构
type TbQuizReviewResponse struct {
	Exercises []Exercise             `json:"exercises"`
	User      UserQuizReviewResponse `json:"user"`
}
type UserQuizReviewResponse struct {
	Uid      uuid.UUID `json:"uid"`
	Username string    `json:"user_name"`
}
