package model

import (
	db "customPro/protoGen/database"
	"github.com/jinzhu/gorm"
)

//protoResponse响应体
type Response struct {
	Id         int32       `form:"id"           json:"id"`
	BaseSet    *BaseSet    `gorm:"ForeignKey:id;AssociationForeignKey:BaseSetId"`
	BaseSetId  int32       `form:"base_set_id"  json:"base_set_id"`
	ResName    string      `form:"res_name"     json:"res_name"`
	ResValue   string      `form:"res_value[]"  json:"res_value"`
	Fields    []map[string]interface{}
}

/**
 * 添加message response响应体
 * @param baseSetId  基础设置主键ID
 * @param resName    message response名称
 * @param resValue   message response字段值
 *
 * @return nil | error
 */
func(res *Response) AddMsgResponse(baseSetId int32, serMethodId int32, resName string, resValue string) error  {
	re := &Response{
		BaseSetId: baseSetId,
		ResName:   resName,
		ResValue:  resValue,
	}

	result := resDB().Create(re)
	if result.Error != nil {
		return result.Error
	}

	//更新服务方法请求参数
	if serMethodId > 0 {
		serM := new(SerMethod).EditRerMethodById(serMethodId, resName, "")
		if serM != nil {
			return serM
		}
	}

	return nil
}

/**
 * 编辑 message response响应体
 *
 * @param id         主键ID
 * @param baseSetId  基础设置主键ID
 * @param resName    message response名称
 * @param resValue   message response字段值
 *
 * @return nil | error
 */
func(res *Response) EditMsgResponse(id int32, baseSetId int32, serMethodId int32, resName string, resValue string) error  {
	re := &Response{
		BaseSetId: baseSetId,
		ResName:   resName,
		ResValue:  resValue,
	}

	result := resDB().Where(&Response{Id:id}).Update(re)
	if result.Error != nil {
		return result.Error
	}

	//更新服务方法请求参数
	if serMethodId > 0 {
		serM := new(SerMethod).EditRerMethodById(serMethodId, resName, "")
		if serM != nil {
			return serM
		}
	}

	return nil
}

/**
 * 根据ID获取一条 message response响应数据
 *
 * @param id         主键ID
 * @return response
 */
func(res *Response) GetOneMsgResponseById(id int32) (*Response, error) {
	re := &Response{Id:id}

	result := resDB().First(re)
	if result.Error != nil {
		return nil, result.Error
	}

	return re, nil
}

/**
 * 根据基础设置主键ID获取 message response响应数据集
 *
 * @param baseSetId    基础设置主键ID
 * @return responses
 */
func(res *Response) GetMsgFromResponsesByBaseSetId(baseSetId int32) ([]*Response, error) {
	var ress []*Response

	result := resDB().Find(ress, &Response{BaseSetId:baseSetId})
	if result.Error != nil {
		return nil, result.Error
	}

	return ress, nil
}


func resDB() *gorm.DB {
	return db.Database().Table("pt_msg_response")
}

