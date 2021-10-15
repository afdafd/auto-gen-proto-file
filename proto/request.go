package proto

import (
	"customPro/protoGen/common"
	"customPro/protoGen/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
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

//获取全部的request列表
func GetAllMsgFromRequests(ctx *gin.Context) {
	results, err := req.GetAllMsgRequests()
	if err != nil {
		common.Error(ctx, err.Error())
		return
	}

	common.Success(ctx, results)
}

//获取请求参数
func getReqParams(ctx *gin.Context) (int32, int32, string, string) {
	tempReq := struct {
		BaseSetId   int32  `json:"base_set_id"`
		ReqName     string `json:"req_name"`
		ReqValue    []map[string]string `form:"req_value[]" json:"req_value"`
		SerMethodId int32  `json:"ser_method_id"`
	}{}

	if err:=ctx.BindJSON(&tempReq); err != nil {
		panic(err.Error())
	}

	reqVls, _ := json.Marshal(tempReq.ReqValue)
	return tempReq.BaseSetId, tempReq.SerMethodId, tempReq.ReqName, string(reqVls)
}
