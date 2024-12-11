package handle

import (
	"wtf-credential/errors"
	"wtf-credential/middleware"
	"wtf-credential/request"
	"wtf-credential/response"
	"wtf-credential/service"

	"github.com/gin-gonic/gin"
)

func GetChapterQuizzes(ctx *gin.Context) {
	req, err := request.BinGetChapterQuizzes(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}

	data, err := service.QuizService.GetChapterQuizzes(ctx, req)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}

func QuizGradeSubmit(ctx *gin.Context) {
	req, err := request.BindGradeSubmit(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	loginUid, _ := middleware.GetUuidFromContext(ctx)
	req.Uid = loginUid

	data, err := service.QuizService.GradeSubmit(ctx, req)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}
