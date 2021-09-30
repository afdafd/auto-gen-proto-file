package proto

import (
	"customPro/protoGen/common"
	"customPro/protoGen/core"
	"customPro/protoGen/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

var baseSet   *model.BaseSet
var service   *model.Service
var request   *model.Request
var response  *model.Response


// 添加proto文件的基础信息
func AddBaseSet(ctx *gin.Context) {
	if err := baseSet.AddBaseSet(getBaseSetReqParams(ctx)); err != nil {
		common.Error(ctx, err.Error())
		return
	}

	common.Success(ctx, "")
}

// 编辑proto文件的基础信息
func EditBaseSet(ctx *gin.Context) {
	_, packageName, className, isAuto := getBaseSetReqParams(ctx)
	err := baseSet.EditBaseSet(getBaseSetId(ctx, true), packageName, className, isAuto)
	if err != nil {
		common.Error(ctx, err.Error())
		return
	}

	common.Success(ctx, gin.H{})
}

// 根据ID获取一条基础设置记录
func GetOneBaseSetById(ctx *gin.Context) {
	result, err := baseSet.GetBaseSetById(getBaseSetId(ctx, false))
	if err != nil {
		common.Error(ctx, err.Error())
		return
	}

	common.Success(ctx, result)
}

// 获取全部基础设置数据集
func GetBaseSetList(ctx *gin.Context) {
	results, err := baseSet.GetBaseSetList()

	if err != nil {
		common.Error(ctx, err.Error())
		return
	}

	common.Success(ctx, results)
}

// 生成 *.proto 文件
func GenerateProtoFileByBaseSet(ctx *gin.Context)  {
	id := getBaseSetId(ctx, true)

	result, _ := baseSet.GetBaseSetById(id)
	sers, _   := service.GetServicesByBaseId(id)
	reqs, _   := request.GetMsgRequestsByBaseId(id)
	ress, _   := response.GetMsgFromResponsesByBaseSetId(id)

	proFile := &core.ProtoFile{
		BaseSet:          result,
		ProtoService:     sers,
		ProtoRequest:     reqs,
		ProtoResResponse: ress,
	}

	//执行生成
	proFile.GenProtoFile()
}

// 获取请求参数
func getBaseSetReqParams(ctx *gin.Context) (int32, string, string, string) {
	setId, _    := strconv.Atoi(ctx.PostForm("set_id"))
	packageName := ctx.PostForm("package_name")
	className   := ctx.PostForm("class_name")
	isAuto      := ctx.PostForm("is_auto_gen_code")

	return int32(setId), packageName, className, isAuto
}

// 获取请求的基础设置表主键ID
func getBaseSetId(ctx *gin.Context, isPost bool) int32 {
	return common.GetId(ctx, "id", isPost)
}
