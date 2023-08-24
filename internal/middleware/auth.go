package middleware

import (
	"minitok/internal/jwt"
	"minitok/internal/unierr"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware Gin middleware: 获取 ctx 中的 token 并检验合法性
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从 query 中取出 token string
		str := ctx.Query("token")

		// 从 multiform 中取出 token string
		if len(str) == 0 {
			str = ctx.PostForm("token")
		}

		//Header取: key:Authorization , value: Bearer <token>
		if len(str) == 0 {
			str = ctx.GetHeader("Authorization")
			if len(str) > 7 && strings.ToLower(str[0:6]) == "bearer" {
				str = str[7:]
			}
		}

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
