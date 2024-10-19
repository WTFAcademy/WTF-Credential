package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseData 数据返回结构
type ResponseData struct {
	Code uint32      `json:"code"`           // 状态码
	Msg  string      `json:"msg"`            // 状态描述
	Data interface{} `json:"data,omitempty"` // 常规json数据或PaginationData数据
}

// JsonSuccess 成功结果返回
// @params ctx  *gin.Context Gin框架上下文
// @params data interface{}  返回数据
func JsonSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, ResponseData{
		Code: 200,
		Data: data,
		Msg:  "success",
	})
	ctx.Abort()
}
