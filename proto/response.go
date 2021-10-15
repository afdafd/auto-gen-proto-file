package proto

import (
	"customPro/protoGen/common"
	"customPro/protoGen/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
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

//获取message response数据集
func GetAllMsgFromResponse(ctx *gin.Context)  {
	results, err := res.GetAllMsgFromResponses()

	if err != nil {
		common.Error(ctx, err.Error())
		return
	}

	common.Success(ctx, results)
}

//获取请求数据
func getResParams(ctx *gin.Context) (int32, int32, string, string) {
	tempRes := struct {
		BaseSetId   int32  `json:"base_set_id"`
		ResName     string `json:"res_name"`
		ResValue    []map[string]string `form:"req_value[]" json:"res_value"`
		SerMethodId int32  `json:"ser_method_id"`
	}{}

	if err:=ctx.BindJSON(&tempRes); err != nil {
		panic(err.Error())
	}

	resVls, _ := json.Marshal(tempRes.ResValue)
	return tempRes.BaseSetId, tempRes.SerMethodId, tempRes.ResName, string(resVls)
}