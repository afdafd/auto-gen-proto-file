package proto

import (
	"customPro/protoGen/common"
	"customPro/protoGen/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

var proSet *model.ProSet


//添加proto项目配置
func AddProSet(ctx *gin.Context) {
	if err := proSet.AddProSet(getProSetParams(ctx)); err != nil {
		common.Error(ctx, err.Error())
		return
	}

	common.Success(ctx, "")
}

//编辑proto项目配置
func EditProSet(ctx *gin.Context) {
	pName, pPath, hName, uName, pwd, id := getProSetParams(ctx)
	err := proSet.EditProSet(id, pName, pPath, hName, uName,pwd)
	if err != nil {
		common.Error(ctx, err.Error())
		return
	}

	common.Success(ctx, "")
}

//获取proto项目配置
func GetProSetById(ctx *gin.Context) {
	result, err := proSet.GetProSetById(common.GetId(ctx, "id", false))

	if err != nil {
		common.Error(ctx, err.Error())
		return
	}

	common.Success(ctx, result)
}

//获取全部proto项目配置
func GetProSets(ctx *gin.Context) {
	results, err := proSet.GetProSets()

	if err != nil {
		common.Error(ctx, err.Error())
		return
	}

	fmt.Println(results)

	common.Success(ctx, results)
}

func GenerateProtoFileByProSet(ctx *gin.Context) {
	id := common.GetId(ctx, "id", true)
	fmt.Println(id)
}

//获取请求参数
func getProSetParams(ctx *gin.Context) (string, string, string, string, string, int32) {
	var pro model.ProSet
	if err := ctx.BindJSON(&pro); err != nil {
		panic(err)
	}

	return pro.ProName, pro.ProPath, pro.HostName, pro.UserName, pro.Pwd, pro.Id
}
