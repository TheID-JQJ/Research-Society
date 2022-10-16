package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//统一结果封装

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, message string) {
	ctx.JSON(httpStatus, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	})
}

func Success(ctx *gin.Context, data gin.H, message string) {
	Response(ctx, http.StatusOK, 200, data, message)
}

func Fail(ctx *gin.Context, message string) {
	Response(ctx, http.StatusOK, 400, gin.H{}, message)
}
