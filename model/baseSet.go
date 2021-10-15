package model

import (
	db "customPro/protoGen/database"
	_ "fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type BaseSets struct {
	Id               int32       `form:"id"                json:"id"`
	ProId            int32       `form:"pro_id"            json:"pro_id"`
	PackageName      string      `form:"package_name"      json:"package_name"`
	ClassName        string      `form:"class_name"        json:"class_name"`
	IsGen            int         `json:"is_gen"`
	IsAutoGenCode    string      `form:"is_auto_gen_code"  json:"is_auto_gen_code"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	ProtoService     []*Service  `json:"proto_service"`
	ProtoRequest     []*Request  `json:"proto_request"`
	ProtoResResponse []*Response `json:"proto_res_response"`
}

/**
 * 添加proto文件的基础设置
 *
 * @param proName      项目名称
 * @param proPath      项目所在的文件路径
 * @param packageName  proto文件里的包名称
 * @param className    proto文件的前缀名称 例：xxx.php
 * @isAutoGenCode      是否根据protoFile自动生成对应的代码；0：否；1：是
 *
 * @return int
 */
func(s *BaseSets) AddBaseSet(setId int32, packageName string, className string, isAutoGenCode string, id int32) error {
	set := &BaseSets{
		ProId:         setId,
		PackageName:   packageName,
		ClassName:     className,
		IsAutoGenCode: isAutoGenCode,
	}

	result := baseDB().Create(set)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

/**
 * 编辑proto文件的基础设置
 *
 * @param id    基础设置表主键ID
 * @param proName      项目名称
 * @param proPath      项目所在的文件路径
 * @param packageName  proto文件里的包名称
 * @param className    proto文件的前缀名称 例：xxx.php
 * @isAutoGenCode      是否根据protoFile自动生成对应的代码；0：否；1：是
 *
 * @return int
 */
func(s *BaseSets) EditBaseSet(setId int32, id int32, packageName string, className string, isAutoGenCode string) error {
	updateValue := &BaseSets{
		PackageName:   packageName,
		ClassName:     className,
		IsAutoGenCode: isAutoGenCode,
		ProId:         setId,
	}

	result := baseDB().Where("id = ?", id).Update(updateValue)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

/**
 * 更新is_gen是否生成proto文件字段
 *
 * @param id      基础设置表主键ID
 * @param isGent  是否生成 0：否；1：是
 *
 * @return error | nil
 */
func (s *BaseSets) UpdateIsGen(id int32, isGen int) error  {
	updateValue := &BaseSets{
		IsGen: isGen,
	}

	if result := baseDB().Where("id = ?", id).Update(updateValue); result.Error != nil {
		return result.Error
	}

	return nil
}

/**
 * 获取一条proto文件的基础设置
 *
 * @param  id  基础设置表主键ID
 * @return *BaseSets
 */
func(s *BaseSets) GetBaseSetById(id int32) (*BaseSets, error) {

	base := &BaseSets{Id: id}
	result := baseDB().First(base)

	if result.Error != nil {
		return nil, result.Error
	}

	return base, nil
}

/**
 * 获取全部proto文件的基础设置数据
 *
 * @return *BaseSets
 */
func(s *BaseSets) GetBaseSetList() ([]*BaseSets, error) {
	var bases []*BaseSets

	if result := baseDB().Find(&bases); result.Error != nil {
		return nil, result.Error
	}

	return bases, nil
}

/**
 * 获取执行项目ID下的全部基础包设置信息集
 *
 * @param ProId 项目ID
 * @return *BaseSets
 */
func(s *BaseSets) GetBaseSetListByProId(proId int32) ([]*BaseSets, error) {
	var bases []*BaseSets

	if result := baseDB().Where("pro_id = ?", proId).Find(&bases); result.Error != nil {
		return nil, result.Error
	}

	return bases, nil
}

/**
 * 通过基础设置ID获取service、request、response
 *
 * @param id 基础设置ID
 * @return []struct
 */
func(s *BaseSets) GetSersAndReqsAndRessByBaseSetId(id int32) ([]*Service, []*Request, []*Response) {
	sers, _ := new(Service).GetServicesByBaseId(id)
	reqs, _ := new(Request).GetMsgRequestsByBaseId(id)
	ress, _ := new(Response).GetMsgFromResponsesByBaseSetId(id)
	return sers, reqs, ress
}


//返回DB实例
func baseDB() *gorm.DB {
	return db.Database().Table("pt_bases_set")
}
