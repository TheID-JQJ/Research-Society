package router

import (
	"hkc.ink/rss/controller"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) {
	//测试
	testGroup := r.Group("/api/test")
	testGroup.GET("/hello", controller.TestHello)
}
