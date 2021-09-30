package proto

import (
	"customPro/protoGen/common"
	"customPro/protoGen/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

var ser *model.Service


//添加proto服务方法名称
func AddService(ctx *gin.Context) {
	setId, serName, serMethods := getParams(ctx)

	if err := ser.AddService(setId, serName, serMethods); err != nil {
		common.Error(ctx, err)
		return
	}

	common.Success(ctx, "")
}

//编辑proto服务方法名
func EditService(ctx *gin.Context)  {
	baseId, serName, serMethods := getParams(ctx)
	err := ser.EditService(common.GetId(ctx, "id", true), baseId, serName, serMethods)

	if err != nil {
		common.Error(ctx, err.Error())
		return
	}

	common.Success(ctx, "")
}

//通过ID获取一条服数据
func GetOneServiceById(ctx *gin.Context)  {
	result, err := ser.GetServiceById(common.GetId(ctx, "id", false))

	if err != nil {
		common.Error(ctx, err.Error())
		return
	}

	common.Success(ctx, result)
}

//通过基础设置ID获取服务信息
func GetServicesByBaseSetId(ctx *gin.Context)  {
	results, err := ser.GetServicesByBaseId(common.GetId(ctx, "base_set_id", false))

	if err != nil {
		common.Error(ctx, err.Error())
		return
	}

	common.Success(ctx, results)
}

//获取参数
func getParams(ctx *gin.Context) (int32, string, []string) {
	setId, _ := strconv.Atoi(ctx.PostForm("base_set_id"))
	serName  := ctx.PostForm("ser_name")
	serMethods := ctx.PostFormArray("ser_methods")

	return int32(setId), serName, serMethods
}
