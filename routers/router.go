package routers

import (
	"customPro/protoGen/proto"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var router *gin.Engine

func init()  {
	r := gin.Default()
	r.Use(Cors())

	r.POST("/user/login", proto.Login)
	r.POST("/user/logout", proto.Logout)
	r.GET("/user/info", proto.UserInfo)

	proSet:= r.Group("/set"); {
		proSet.POST("/add-pro-set", proto.AddProSet)
		proSet.PUT("/edit-pro-set", proto.EditProSet)
		proSet.GET("/get-one-pro-set-by-id", proto.GetProSetById)
		proSet.GET("/get-pro-set-list", proto.GetProSets)
		proSet.POST("/gen-proto-file", proto.GenerateProtoFileByProSet)
	}

	base:= r.Group("/base"); {
		base.POST("/add-proto-base-set", proto.AddBaseSet)
		base.PUT("/edit-proto-base-set", proto.EditBaseSet)
		base.GET("/get-one-base-by-id", proto.GetOneBaseSetById)
		base.GET("/get-base-list", proto.GetBaseSetList)
		base.POST("/gen-proto-file", proto.GenerateProtoFileByBaseSet)
		base.GET("/gen-all-by-id", proto.GetSersAndReqsAndRessByBaseSetId)
	}

	ser := r.Group("/ser"); {
		ser.POST("/add-proto-service", proto.AddService)
		ser.PUT("/edit-proto-service", proto.EditService)
		ser.GET("get-one-service-by-id", proto.GetOneServiceById)
		ser.GET("get-service-list", proto.GetServicesByBaseSetId)
		ser.GET("get-all-service-list", proto.GetAllServicesList)
	}

	req := r.Group("/req"); {
		req.POST("/add-proto-msg-request", proto.AddMsgFromRequest)
		req.PUT("/edit-proto-msg-request", proto.EditMsgFromRequest)
		req.GET("/get-one-req-by-id", proto.GetOneMsgFromRequestById)
		req.GET("/get-reqs-by-set-id", proto.GetMsgFromRequestsByBaseSetId)
	}

	res := r.Group("/res"); {
		res.POST("/add-proto-msg-response", proto.AddMsgFromResponse)
		res.PUT("/edit-proto-msg-response", proto.EditMsgFromResponse)
		res.GET("/get-one-res-by-id", proto.GetOneMsgFromResponseById)
		res.GET("/get-ress-by-set-id", proto.GetMsgFromResponsesByBaseSetId)
	}

	router = r
}

func Run()  {
	//gin.SetMode(gin.DebugMode)
	router.Run(":9999")
}

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		origin := ctx.Request.Header.Get("Origin")  //请求头部
		fmt.Println("执行中间件代码", origin)

		//ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		//服务器支持的所有跨域请求的方法
		ctx.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE,PATCH,UPDATE")

		//允许跨域设置可以返回其他子段，可以自定义字段
		ctx.Header("Access-Control-Allow-Headers", "Origin,Accept,Authorization,Content-Length,Content-Type,Token,x-token,session,Access-Control-Allow-Origin")

		// 允许浏览器（客户端）可以解析的头部
		ctx.Header("Access-Control-Expose-Headers", "Content-Length,Content-Type,Access-Control-Allow-Origin,Access-Control-Allow-Headers")

		//设置缓存时间
		//ctx.Header("Access-Control-Max-Age", "172800")

		//允许客户端传递校验信息比如 cookie
		ctx.Header("Access-Control-Allow-Credentials", "true")

		//允许类型校验
		if method := ctx.Request.Method; method == "OPTIONS" {
			//ctx.JSON(http.StatusOK, "ok!")
			ctx.AbortWithStatus(http.StatusNoContent)
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		ctx.Next()
	}
}
