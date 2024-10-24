package handle

import (
	"github.com/gin-gonic/gin"
	"wtf-credential/errors"
	"wtf-credential/request"
	"wtf-credential/response"
	"wtf-credential/service"
)

func GetAllCourse(ctx *gin.Context) {

	req, err := request.BinGetAllCourse(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	data, err := service.GetAllCourse(ctx, req)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}

func GetCourseInfo(ctx *gin.Context) {
	req, err := request.BinGetCourseInfo(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}

	data, err := service.GetCourseInfo(ctx, req)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}

	response.JsonSuccess(ctx, data)
}
