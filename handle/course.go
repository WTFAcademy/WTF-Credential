package handle

import (
	"github.com/gin-gonic/gin"
	"wtf-credential/errors"
	"wtf-credential/middleware"
	"wtf-credential/request"
	"wtf-credential/response"
	"wtf-credential/service"
)

func GetAllCourse(ctx *gin.Context) {

	data, err := service.GetAllCourse(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}

func GetCoursesByType(ctx *gin.Context) {
	data, err := service.GetCoursesByType(ctx)
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

func GetCourseQuizzes(ctx *gin.Context) {
	loginUid, ok := middleware.GetUuidFromContext(ctx)
	if !ok {
		ctx.JSON(200, errors.Entity("failed to retrieve user from context"))
		return
	}
	req, err := request.BinGetCourseQuizzes(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	data, err := service.GetCourseQuizzes(ctx, req, loginUid)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}

func GetUserCourseLesson(ctx *gin.Context) {
	loginUid, ok := middleware.GetUuidFromContext(ctx)
	if !ok {
		ctx.JSON(200, errors.Entity("failed to retrieve user from context"))
		return
	}

	req, err := request.BinGetUserCourseLesson(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	data, err := service.GetUserCourseLesson(ctx, req, loginUid)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}

func GetStatistics(ctx *gin.Context) {
	data, err := service.GetStatistics(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}

func GetCourseByPath(ctx *gin.Context) {
	req, err := request.BinGetCourseByPath(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}

	data, err := service.GetCourseByPath(ctx, req)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}

	response.JsonSuccess(ctx, data)

}

func GetCourseChapters(ctx *gin.Context) {
	loginUid := middleware.GetCourseChaptersUidFromContext(ctx) // 忽略错误处理
	req, err := request.BinGetCourseChaptersByPath(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	var data []response.GetCourseChapters
	if loginUid != "" {
		data, err = service.GetUserCourseChaptersByPath(ctx, req, loginUid)
		if err != nil {
			ctx.JSON(200, errors.Unknown(err))
			return
		}
	} else {
		data, err = service.GetCourseChaptersByPath(ctx, req)
		if err != nil {
			ctx.JSON(200, errors.Unknown(err))
			return
		}
	}
	response.JsonSuccess(ctx, data)
}

func GetChapterDetailsByID(ctx *gin.Context) {
	loginUid := middleware.GetCourseChaptersUidFromContext(ctx) // 忽略错误处理
	req, err := request.BinGetChapterDetailsByID(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	var data response.GetChapterDetailsByID
	if loginUid != "" {
		data, err = service.GetUserGetChapterDetailsByID(ctx, req, loginUid)
		if err != nil {
			ctx.JSON(200, errors.Unknown(err))
			return
		}
	} else {
		data, err = service.GetChapterDetailsByID(ctx, req)
		if err != nil {
			ctx.JSON(200, errors.Unknown(err))
			return
		}
	}
	response.JsonSuccess(ctx, data)
}
