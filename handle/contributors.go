package handle

import (
	"github.com/gin-gonic/gin"
	"wtf-credential/errors"
	"wtf-credential/request"
	"wtf-credential/response"
	"wtf-credential/service"
)

// ContributorsList 按不同的键返回所有 Redis 中的贡献者信息
func GetContributorsList(ctx *gin.Context) {
	var getContributorsListRequest request.GetContributorsList
	if err := ctx.ShouldBindJSON(&getContributorsListRequest); err != nil {
		ctx.JSON(200, errors.Entity("param error"))
		return
	}

	data, err := service.GetContributorsList(ctx, getContributorsListRequest)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}
