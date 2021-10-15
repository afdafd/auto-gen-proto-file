package model

import (
	db "customPro/protoGen/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//protoRpc服务体
type Service struct {
	Id         int32         `form:"id" json:"id"`
	BaseSetId  int32         `form:"base_set_id" json:"base_set_id"`
	SerName    string        `form:"ser_name" json:"ser_name"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
	//BaseSets    *BaseSets      `gorm:"ForeignKey:id;AssociationForeignKey:BaseSetId"`
	SerMethods []*SerMethod   `json:"ser_methods"`
}

//protoRpc服务方法
type SerMethod struct {
	Id            int32     `form:"id" json:"id"`
	SerId         int32     `form:"ser_id"        json:"ser_id"`
	MethodName    string    `form:"method_name"   json:"method_name"`
	RequestName   string    `form:"request_name"  json:"request_name"`
	ResponseName  string    `form:"response_name" json:"response_name"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	Service       *Service  `gorm:"ForeignKey:id;AssociationForeignKey:SerId"`
}


/**
 * 添加proto服务名
 *
 * @param baseSetId  基础设置表主键ID
 * @param serName    服务名
 *
 * @return nil | error
 */
func(s *Service) AddService(baseSetId int32, serName string, serMethods []string, id int32) error {
	ser := &Service{
		BaseSetId:      baseSetId,
		SerName:        serName,
	}

	serDB().Begin()
	result := serDB().Create(ser)
	if result.Error != nil {
		serDB().Rollback()
		return result.Error
	}

	for _, v := range serMethods {
		sm := &SerMethod{
			SerId:       ser.Id,
			MethodName:  v,
		}

		resM := serMethodDB().Create(sm)
		if resM.Error != nil {
			serDB().Rollback()
			return resM.Error
		}
	}


	serDB().Commit()
	return nil
}

/**
 * 编辑proto服务名
 *
 * @param id         服务表主键ID
 * @param baseSetId  基础设置表主键ID
 * @param serName    服务名
 *
 * @return nil | error
 */
func(s *Service) EditService(id int32, baseId int32, serName string, serMethods []string) error {
	ser := &Service{
		BaseSetId: baseId,
		SerName:   serName,
	}

	result := serDB().Where("id = ?", id).Update(ser)
	if result.Error != nil {
		return result.Error
	}

	//先删除
	serMDeRes := serMethodDB().Where("ser_id = ?", id).Delete(&SerMethod{})
	if serMDeRes.Error != nil {
		return serMDeRes.Error
	}

	for _, v := range serMethods {
		sm := &SerMethod{
			SerId:      id,
			MethodName: v,
		}

		addRes := serMethodDB().Create(sm)
		if addRes.Error != nil {
			return addRes.Error
		}
	}

	return nil
}

/**
 * 编辑proto服务方法参数值
 *
 * @param id         服务表主键ID
 * @param baseSetId  基础设置表主键ID
 * @param serName    服务名
 *
 * @return nil | error
 */
func(sm *SerMethod) EditRerMethodById (id int32, reqName string, resName string) error {
	seM:= &SerMethod{
		RequestName:  reqName,
		ResponseName: resName,
	}

	result := serMethodDB().Where("id = ?", id).Update(&seM)
	if result.Error != nil {
		return result.Error
	}

	return nil
}


/**
 * 根据基础设置表主键ID获取 proto服务 数据信息
 *
 * @param baseSetId  基础设置表主键ID
 * @return []*Service | error
 */
func(s *Service) GetServicesByBaseId(baseSetId int32) ([]*Service, error) {
	var sers []*Service

	results := serDB().Where("base_set_id = ?", baseSetId).Find(&sers)
	if results.Error != nil {
		return nil, results.Error
	}

	var seMs []*SerMethod

	for _, v := range sers {
		if methods := serMethodDB().Where("ser_id = ?", v.Id).Find(&seMs); methods.Error != nil {
			return nil, methods.Error
		}

		for _, mv := range seMs {
			v.SerMethods = append(v.SerMethods, mv)
		}
	}

	return sers, nil
}

/**
 * 获取 全部的proto服务 数据信息
 *
 * @return []*Service | error
 */
func(s *Service) GetAllServiceList() ([]*Service, error) {
	var sers []*Service

	results := serDB().Find(&sers)
	if results.Error != nil {
		return nil, results.Error
	}

	var seMs []*SerMethod

	for _, v := range sers {
		if methods := serMethodDB().Where("ser_id = ?", v.Id).Find(&seMs); methods.Error != nil {
			return nil, methods.Error
		}

		for _, mv := range seMs {
			v.SerMethods = append(v.SerMethods, mv)
		}
	}

	return sers, nil
}


/**
 * 根据ID获取一条proto服务数据信息
 *
 * @param Id  基础设置表主键ID
 * @return *Service | error
 */
func(s *Service) GetServiceById(id int32) (*Service, error)  {
	ser := &Service{Id:id}

	result := serDB().First(ser)
	if result.Error != nil {
		return nil, result.Error
	}

	var seMs []*SerMethod
	if methods := serMethodDB().Where("ser_id = ?", id).Find(&seMs); methods.Error != nil {
		return nil, methods.Error
	}

	for _, mv := range seMs {
		ser.SerMethods = append(ser.SerMethods, mv)
	}

	return ser, nil
}

//返回DB实例
func serDB() *gorm.DB {
	return db.Database().Table("pt_service")
}

func serMethodDB() *gorm.DB {
	return db.Database().Table("pt_ser_method")
}