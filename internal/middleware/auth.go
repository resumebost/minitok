package middleware

import (
	"github.com/gin-gonic/gin"
	"minitok/internal/jwt"
	"minitok/internal/unierr"
	"net/http"
)

// AuthMiddleware Gin middleware: 获取 ctx 中的 token 并检验合法性
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 取出 token string
		str := ctx.Query("token")

		if len(str) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, unierr.NoTokenError)
			return
		}

		claims, err := jwt.ParseToken(str)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, unierr.InvalidTokenError)
			return
		}

		// 写入信息
		ctx.Set("username", claims.Username)
		ctx.Set("id", claims.ID)
		ctx.Next()
	}
}
