package cmd

import (
	"hkc.ink/rss/database"
	"hkc.ink/rss/router"

	"github.com/gin-gonic/gin"
)

func Execute() {
	//gorm
	db := database.InitDB()
	defer database.CloseDB(db)

	//gin
	r := gin.Default()
	router.CollectRouter(r)
	r.Run(":8080")
}
