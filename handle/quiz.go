package handle

import (
	"fmt"
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
	// 1. 获取登录用户的 UID
	loginUid, ok := middleware.GetUuidFromContext(ctx)
	if !ok {
		ctx.JSON(200, errors.Entity("failed to retrieve user from context"))
		return
	}
	fmt.Println("Successfully retrieved loginUid:", loginUid)

	// 2. 绑定请求数据
	req, err := request.BindGradeSubmit(ctx)
	if err != nil {
		fmt.Println("Failed to bind grade submit request:", err)
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	fmt.Println("Successfully bound request:", req)

	// 3. 设置请求数据中的 UID
	req.Uid = loginUid
	fmt.Println("Set req.Uid to:", req.Uid)

	// 4. 调用服务层处理成绩提交
	data, err := service.QuizService.GradeSubmit(ctx, req)
	if err != nil {
		fmt.Println("Error occurred in GradeSubmit:", err)
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	fmt.Println("GradeSubmit successful, response data:", data)

	// 5. 返回成功响应
	response.JsonSuccess(ctx, data)
}
