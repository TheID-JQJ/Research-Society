package router

import (
	"hkc.ink/rss/controller"
	"hkc.ink/rss/middleware"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "hkc.ink/rss/docs"
)

func CollectRouter(r *gin.Engine) {
	//测试
	testGroup := r.Group("/api/test")
	testGroup.GET("/hello", middleware.AuthMiddleware(), middleware.CasbinHandler(), controller.TestHello)
	testGroup.POST("/register", controller.TestRegister)
	testGroup.POST("/insert", controller.TestInsert)
	testGroup.POST("/delete/:id", controller.TestDelete)
	testGroup.POST("/update", controller.TestUpdate)
	testGroup.GET("/get/:id", controller.TestGet)
	testGroup.GET("/get/all", controller.TestGetAll)

	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
