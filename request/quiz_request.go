package request

import "github.com/gin-gonic/gin"

type GetChapterQuizzes struct {
	Path     string `uri:"path" binding:"required"`         // 使用 uri 标签绑定路径参数
	RothPath string `uri:"chapter_path" binding:"required"` // 使用 uri 标签绑定路径参数
}

func BinGetChapterQuizzes(c *gin.Context) (*GetChapterQuizzes, error) {
	var req GetChapterQuizzes
	if err := c.ShouldBindUri(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
