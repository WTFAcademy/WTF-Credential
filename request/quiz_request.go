package request

import (
	"github.com/gin-gonic/gin"
)

type GetChapterQuizzes struct {
	Path      string `uri:"course_path" binding:"required"`  // 使用 uri 标签绑定路径参数
	RoutePath string `uri:"chapter_path" binding:"required"` // 使用 uri 标签绑定路径参数
}

func BinGetChapterQuizzes(c *gin.Context) (*GetChapterQuizzes, error) {
	var req GetChapterQuizzes
	if err := c.ShouldBindUri(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func BindGradeSubmit(c *gin.Context) (*QuizGradeSubmitReq, error) {
	var req QuizGradeSubmitReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

type QuizGradeSubmitReq struct {
	Uid       string          `json:"uid"`
	ChapterId int64           `json:"chapter_id"`
	CourseId  int64           `json:"course_id"`
	Answers   []AnswerRequest `json:"answers"`
}

type AnswerRequest struct {
	Id      int64    `json:"id"`
	Answers []string `json:"answers"`
}
