package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"wtf-credential/configs"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)

type MyClaims struct {
	email string `json:"email"`
	jwt.RegisteredClaims
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		mc, err := ParseToken(c)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "Please login first",
			})
			c.Abort()
			return
		}
		// Save the currently requested username information to the requested context c
		c.Set("login_uid", mc.Subject)
		c.Next() // Subsequent processing functions can use c.Get("username") to obtain the currently requested user information
	}
}

func CourseJWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取 token
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			// 如果没有 token，继续处理请求，不返回错误
			c.Next()
			return
		}

		// 如果有 token，尝试解析
		mc, err := ParseToken(c)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "Please login first",
			})
			c.Abort()
			return
		}

		// 保存当前请求的用户 ID 到上下文中
		c.Set("login_uid", mc.Subject)
		c.Next() // 继续后续处理
	}
}

func GetLoginUid(c *gin.Context) string {
	if c.Query("network") == "stark_net" && c.Query("login_uid") != "" {
		return c.Query("login_uid")
	}
	mc, err := ParseToken(c)
	if err != nil {
		return ""
	}
	return mc.Subject
}

func GetUuidFromContext(c *gin.Context) (string, bool) {
	uuid, exists := c.Get("login_uid")
	if !exists {
		// 如果找不到 login_uid，返回 401 错误并退出
		c.JSON(http.StatusUnauthorized, gin.H{"message": "无法获取id"})
		return "", false
	}

	uuidStr, ok := uuid.(string)
	if !ok {
		// 如果转换为 string 类型失败，返回 500 错误并退出
		c.JSON(http.StatusInternalServerError, gin.H{"message": "id格式错误"})
		return "", false
	}

	return uuidStr, true
}

func GetCourseChaptersUidFromContext(c *gin.Context) string {
	uuid, exists := c.Get("login_uid")
	if !exists {
		// 返回空字符串
		return ""
	}

	uuidStr, ok := uuid.(string)
	if !ok {
		// 返回空字符串
		return ""
	}

	return uuidStr
}

func ParseToken(c *gin.Context) (*MyClaims, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("invalid Authorization")
	}

	// 直接使用整个 authHeader 作为 token
	token, err := jwt.ParseWithClaims(authHeader, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(configs.Config().JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
