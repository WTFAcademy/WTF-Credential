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
	loginUid, ok := middleware.GetUuidFromContext(ctx)
	if !ok {
		ctx.JSON(200, errors.Entity("failed to retrieve user from context"))
		return
	}
	req, err := request.BinGetChapterQuizzes(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}

	data, err = service.GetChapterQuizzes(ctx, req)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}
