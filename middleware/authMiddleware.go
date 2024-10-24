package middleware

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"wtf-credential/configs"

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
		c.JSON(http.StatusUnauthorized, gin.H{"message": "无法获取id"})
		return "", false
	}

	uuidStr, ok := uuid.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "id格式错误"})
		return "", false
	}

	return uuidStr, true
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
