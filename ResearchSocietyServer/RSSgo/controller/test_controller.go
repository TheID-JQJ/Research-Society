package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"hkc.ink/rss/database"
	"hkc.ink/rss/dto"
	"hkc.ink/rss/response"
	"hkc.ink/rss/service"
)

var userService service.UserService

func TestHello(ctx *gin.Context) {
	response.Success(ctx, gin.H{}, "hello hkc")
}

func TestInsert(ctx *gin.Context) {
	//接收数据
	rm := dto.RegisterModel{}
	if err := ctx.ShouldBindJSON(&rm); err != nil {
		log.Println("this is insert:" + err.Error())
	}

	//写入数据库并返回结果
	db := database.GetDB()
	message := ""
	if err := userService.Register(db, &rm); err != nil {
		message = err.Error()
	} else {
		message = "注册成功"
	}

	response.Success(ctx, gin.H{}, message)
}
