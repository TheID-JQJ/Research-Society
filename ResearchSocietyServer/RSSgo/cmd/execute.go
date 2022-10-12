package cmd

import (
	"hkc.ink/rss/router"

	"github.com/gin-gonic/gin"
)

func Execute() {
	//gin
	r := gin.Default()
	router.CollectRouter(r)
	r.Run(":8080")
}
