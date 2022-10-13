package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"hkc.ink/rss/database"
	"hkc.ink/rss/response"
	"hkc.ink/rss/service"
	"hkc.ink/rss/utils"
)

var jwtUtil utils.JwtUtil
var userService service.UserService

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取请求头Authorization
		authorization := ctx.GetHeader("Authorization")
		// log.Println(authorization)

		//校验
		if authorization == "" || !strings.HasPrefix(authorization, "Bearer") {
			response.Response(ctx, http.StatusUnauthorized, 401, gin.H{}, "权限不足")
			ctx.Abort()
			return
		}
		authorization = authorization[7:]
		// log.Println(authorization)

		//解析
		token, claims, err := jwtUtil.ParseToken(authorization)
		if err != nil || !token.Valid {
			response.Response(ctx, http.StatusUnauthorized, 401, gin.H{}, "权限不足")
			ctx.Abort()
			return
		}

		//验证userID
		user, idErr := userService.GetByID(database.GetDB(), claims.UserID)
		if idErr != nil || user.ID <= 0 {
			response.Response(ctx, http.StatusUnauthorized, 401, gin.H{}, "权限不足")
			ctx.Abort()
			return
		}

		//将用户信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}
