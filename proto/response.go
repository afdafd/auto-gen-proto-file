package proto

import (
	"customPro/protoGen/common"
	"customPro/protoGen/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

var res *model.Response


//添加message response响应体
func AddMsgFromResponse(ctx *gin.Context) {
	result := res.AddMsgResponse(getResParams(ctx))
	if result != nil {
		common.Error(ctx, result.Error())
		return
	}

	common.Success(ctx, "")
}

//编辑message response响应体
func EditMsgFromResponse(ctx *gin.Context)  {
	id := common.GetId(ctx, "id", true)
	setId, serMethodId, resName, resValue := getResParams(ctx)

	result := res.EditMsgResponse(id, setId, serMethodId, resName, resValue)
	if result != nil {
		common.Error(ctx, result.Error())
		return
	}

	common.Success(ctx, "")
}

//根据主键ID获取一条message response响应体
func GetOneMsgFromResponseById(ctx *gin.Context)  {
	id := common.GetId(ctx, "id", false)
	result, err := res.GetOneMsgResponseById(id)

	if err != nil {
		common.Error(ctx, err.Error())
		return
	}

	common.Success(ctx, result)
}

//根据基础设置ID获取message response数据集
func GetMsgFromResponsesByBaseSetId(ctx *gin.Context)  {
	baseSetId := common.GetId(ctx, "base_set_id", false)
	results, err := res.GetMsgFromResponsesByBaseSetId(baseSetId)

	if err != nil {
		common.Error(ctx, err.Error())
		return
	}

	common.Success(ctx, results)
}

//获取请求数据
func getResParams(ctx *gin.Context) (int32, int32, string, string) {
	setId, _ := strconv.Atoi(ctx.PostForm("base_set_id"))
	serMethodId, _ := strconv.Atoi(ctx.PostForm("ser_method_id"))
	resName  := ctx.PostForm("res_name")
	resValue := ctx.PostForm("res_value")

	return int32(setId), int32(serMethodId), resName, resValue
}