package common

import "github.com/gin-gonic/gin"

func Success(ctx *gin.Context, data interface{})  {
	ctx.JSON(200, gin.H{
		"code": 1,
		"data": data,
	})
}

func Error(ctx *gin.Context, data interface{})  {
	ctx.JSON(200, gin.H{
		"code": -1,
		"data": data,
	})
}
