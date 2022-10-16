package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hkc.ink/rss/database"
	"hkc.ink/rss/model"
	"hkc.ink/rss/response"
	"hkc.ink/rss/utils"
)

// var casbinService service.CasbinService
var casbinUtil utils.CasbinUtil

func CasbinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取请求path
		obj := ctx.Request.URL.Path
		//获取请求方法
		act := ctx.Request.Method
		//角色
		value, exists := ctx.Get("user")
		if !exists {
			response.Response(ctx, http.StatusUnauthorized, 401, gin.H{}, "权限不足")
			ctx.Abort()
			return
		}
		sub := (value.(model.User)).Email

		//新增几个策略
		//e.AddPolicy("admin", "/v1/customer/customer", "GET")
		////删除策略
		//e.RemovePolicy("admin", "/api/v1/hello", "GET")
		////获取策略
		//list := e.GetPolicy()
		//for _, vlist := range list {
		// for _, v := range vlist {
		//    fmt.Printf("value: %s, ", v)
		// }
		//}

		//判断策略是否存在
		success := casbinUtil.Enforce(database.GetDB(), sub, obj, act)
		if success {
			ctx.Next()
		} else {
			response.Response(ctx, http.StatusUnauthorized, 401, gin.H{}, "权限不足")
			ctx.Abort()
			return
		}
	}
}
