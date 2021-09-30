package proto

import (
	"customPro/protoGen/common"
	"customPro/protoGen/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

var baseSet *model.BaseSet


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
	res, err := baseSet.GetBaseSetById(getBaseSetId(ctx, false))
	if err != nil {
		common.Error(ctx, err.Error())
		return
	}

	common.Success(ctx, res)
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

// 生成 *.proto文件
func GenerateProtoFile(ctx *gin.Context)  {
	baseSetId := getBaseSetId(ctx, true)
	fmt.Println(baseSetId)
	//...
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
