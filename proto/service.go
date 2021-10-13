package proto

import (
	"customPro/protoGen/common"
	"customPro/protoGen/model"
	"github.com/gin-gonic/gin"
)

var ser *model.Service


//添加proto服务方法名称
func AddService(ctx *gin.Context) {
	if err := ser.AddService(getParams(ctx)); err != nil {
		common.Error(ctx, err)
		return
	}

	common.Success(ctx, "")
}

//编辑proto服务方法名
func EditService(ctx *gin.Context)  {
	baseId, serName, serMethods, id := getParams(ctx)
	err := ser.EditService(id, baseId, serName, serMethods)

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

//获取全部的服务信息
func GetAllServicesList(ctx *gin.Context)  {
	results, err := ser.GetAllServiceList()

	if err != nil {
		common.Error(ctx, err.Error())
		return
	}

	common.Success(ctx, results)
}

//获取参数
func getParams(ctx *gin.Context) (int32, string, []string, int32) {
	var temp struct {
		Id        int32   `json:"id"`
		SerName   string `json:"ser_name"`
		BaseSetId int32 `json:"base_set_id"`
		SerMethods []string `json:"ser_methods"`
	}

	if err := ctx.BindJSON(&temp); err != nil {
		panic(err)
	}

	return temp.BaseSetId, temp.SerName, temp.SerMethods, temp.Id
}
