package proto

import (
	"customPro/protoGen/common"
	"customPro/protoGen/core"
	"customPro/protoGen/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var baseSet   *model.BaseSets
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
	setId, packageName, className, isAuto, id := getBaseSetReqParams(ctx)
	err := baseSet.EditBaseSet(setId, id, packageName, className, isAuto)
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

// 获取执行项目ID下的全部基础包设置信息集
func GetBaseSetListByProId(ctx *gin.Context) {
	 proId := common.GetId(ctx, "pro_id", false)
	 results, err := baseSet.GetBaseSetListByProId(proId)

	 if err != nil {
	 	common.Error(ctx, err.Error())
	 	return
	 }

	 common.Success(ctx, results)
}

//通过基础设置ID获取service、request、response
func GetSersAndReqsAndRessByBaseSetId(ctx *gin.Context)  {
	sers, reqs, ress := baseSet.GetSersAndReqsAndRessByBaseSetId(getBaseSetId(ctx, false))

	common.Success(ctx, gin.H{
		"sers": sers,
		"reqs": reqs,
		"ress": ress,
	})
}

// 生成 *.proto 文件
func GenerateProtoFileByBaseSet(ctx *gin.Context)  {
	id := getBaseSetId(ctx, false)

	fmt.Println("ID", id)

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

	//更新is_gen字段值
	if err := baseSet.UpdateIsGen(id, 1); err != nil {
		log.Fatal("更新is_gen字段失败")
	}

	common.Success(ctx, "")
}

// 获取请求参数
func getBaseSetReqParams(ctx *gin.Context) (int32, string, string, string, int32) {
	var baseSet model.BaseSets
	if err := ctx.BindJSON(&baseSet); err != nil {
		panic(err)
	}

	return baseSet.ProId, baseSet.PackageName, baseSet.ClassName, baseSet.IsAutoGenCode, baseSet.Id
}

// 获取请求的基础设置表主键ID
func getBaseSetId(ctx *gin.Context, isPost bool) int32 {
	return common.GetId(ctx, "id", isPost)
}
