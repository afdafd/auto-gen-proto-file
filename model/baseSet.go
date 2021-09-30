package model

import (
	db "customPro/protoGen/database"
	"fmt"
	_"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type BaseSet struct {
	Id            int32     `form:"id"                json:"id"`
	SetId         int32     `form:"set_id"            json:"set_id"`
	PackageName   string    `form:"package_name"      json:"package_name"`
	ClassName     string    `form:"class_name"        json:"class_name"`
	IsAutoGenCode string    `form:"is_auto_gen_code"  json:"is_auto_gen_code"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	ProtoService      []*Service   `json:"proto_service"`
	ProtoRequest      []*Request   `json:"proto_request"`
	ProtoResResponse  []*Response  `json:"proto_res_response"`
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
func(s *BaseSet) AddBaseSet(setId int32, packageName string, className string, isAutoGenCode string) error {
	set := &BaseSet{
		SetId:setId,
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
func(s *BaseSet) EditBaseSet(id int32, packageName string, className string, isAutoGenCode string) error {
	updateValue := &BaseSet{
		PackageName:   packageName,
		ClassName:     className,
		IsAutoGenCode: isAutoGenCode,
	}

	fmt.Println("自动：",isAutoGenCode)
	result := baseDB().Where("id = ?", id).Update(updateValue)

	if result.Error != nil {
		panic(result.Error)
	}

	return nil
}

/**
 * 获取一条proto文件的基础设置
 *
 * @param  id  基础设置表主键ID
 * @return *BaseSet
 */
func(s *BaseSet) GetBaseSetById(id int32) (*BaseSet, error) {

	base := &BaseSet{Id: id}
	result := baseDB().First(base)

	if result.Error != nil {
		return nil, result.Error
	}

	return base, nil
}

/**
 * 获取全部proto文件的基础设置数据
 *
 * @return *BaseSet
 */
func(s *BaseSet) GetBaseSetList() ([]*BaseSet, error) {
	var bases []*BaseSet

	if result := baseDB().Find(&bases); result.Error != nil {
		return nil, result.Error
	}

	return bases, nil
}


//返回DB实例
func baseDB() *gorm.DB {
	return db.Database().Table("pt_base_set")
}
