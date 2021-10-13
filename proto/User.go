package proto

import (
	"customPro/protoGen/common"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context)  {
	common.Success(ctx, "abcd123456")
}

func Logout(ctx *gin.Context)  {
	common.Success(ctx, "success")
}

func UserInfo(ctx *gin.Context)  {
	common.Success(ctx, gin.H{
		"avatar": "http://static.runoob.com/images/demo/demo2.jpg",
		"name": "admin",
	})
}
