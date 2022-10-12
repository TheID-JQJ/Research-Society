package controller

import "github.com/gin-gonic/gin"

func TestHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello hkc",
	})
}
