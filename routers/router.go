package routers

import (
	"customPro/protoGen/proto"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init()  {
	r := gin.Default()

	proSet:= r.Group("/set"); {
		proSet.POST("/add-pro-set", proto.AddProSet)
		proSet.PUT("/edit-pro-set", proto.EditProSet)
		proSet.GET("/get-one-pro-set-by-id", proto.GetProSetById)
		proSet.GET("/get-pro-set-list", proto.GetProSets)
	}

	base:= r.Group("/base"); {
		base.POST("/add-proto-base-set", proto.AddBaseSet)
		base.PUT("/edit-proto-base-set", proto.EditBaseSet)
		base.GET("/get-one-base-by-id", proto.GetOneBaseSetById)
		base.GET("/get-base-list", proto.GetBaseSetList)
		base.POST("/gen-code", proto.GenerateProtoFile)
	}

	ser := r.Group("/ser"); {
		ser.POST("/add-proto-service", proto.AddService)
		ser.PUT("/edit-proto-service", proto.EditService)
		ser.GET("get-one-service-by-id", proto.GetOneServiceById)
		ser.GET("get-service-list", proto.GetServicesByBaseSetId)
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
	gin.SetMode(gin.DebugMode)
	router.Run(":9999")
}