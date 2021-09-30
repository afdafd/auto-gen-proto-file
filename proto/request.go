package proto

import (
	"customPro/protoGen/common"
	"customPro/protoGen/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

var req *model.Request


//添加message request
func AddMsgFromRequest(ctx *gin.Context) {
	result := req.AddMsgRequest(getReqParams(ctx))
	if result !=nil {
		common.Error(ctx, result.Error())
		return
	}

	common.Success(ctx, "")
}

//编辑message request
func EditMsgFromRequest(ctx *gin.Context) {
	id := common.GetId(ctx, "id", true)
	baseSetId, resMethodId, reqName, reqValue := getReqParams(ctx)

	result := req.EditMsgRequest(id, baseSetId, resMethodId, reqName, reqValue)
	if result != nil {
		common.Error(ctx, result.Error())
		return
	}

	common.Success(ctx, "")

}

//根据ID获取一条message request
func GetOneMsgFromRequestById(ctx *gin.Context)  {
	result, err := req.GetOneMsgRequestById(common.GetId(ctx, "id", false))
	if err != nil {
		common.Error(ctx,err.Error())
		return
	}

	common.Success(ctx, result)
}

//通过基础设置ID获取message request 数据集
func GetMsgFromRequestsByBaseSetId(ctx *gin.Context) {
	results, err := req.GetMsgRequestsByBaseId(common.GetId(ctx, "base_set_id", false))
	if err != nil {
		common.Error(ctx, err.Error())
		return
	}

	common.Success(ctx, results)
}

//获取请求参数
func getReqParams(ctx *gin.Context) (int32, int32, string, string) {
	baseSetId, _   := strconv.Atoi(ctx.PostForm("base_set_id"))
	serMethodId, _ := strconv.Atoi(ctx.PostForm("ser_method_id"))
	reqName  := ctx.PostForm("req_name")
	reqValue := ctx.PostForm("req_value")

	return int32(baseSetId), int32(serMethodId), reqName, reqValue
}
