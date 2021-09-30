package model

import (
	db "customPro/protoGen/database"
	"github.com/jinzhu/gorm"
)

//protoMessage请求体
type Request struct {
	Id         int32       `form:"id"           json:"id"`
	BaseSet    *BaseSet    `gorm:"ForeignKey:id;AssociationForeignKey:BaseSetId"`
	BaseSetId  int32       `form:"base_set_id"  json:"base_set_id"`
	ReqName    string      `form:"req_name"     json:"req_name"`
	ReqValue   string      `form:"req_value[]"  json:"req_value"`
	Fields     []map[string]interface{}
}

/**
 * 添加 message request 请求体信息
 *
 * @param baseSetId    基础设置ID
 * @param reqName      request名称
 * @param reqValue     request字段值
 *
 * @return error | nil
 */
func(req *Request) AddMsgRequest(baseSetId int32, serMethodId int32, reqName string, reqValue string) error  {
	reqs := &Request{
		BaseSetId: baseSetId,
		ReqName:   reqName,
		ReqValue:  reqValue,
	}

	result := reqDB().Create(reqs)
	if result.Error != nil {
		return result.Error
	}

	//更新服务方法请求参数
	if serMethodId > 0 {
		serM := new(SerMethod).EditRerMethodById(serMethodId, reqName, "")
		if serM != nil {
			return serM
		}
	}

	return nil
}

/**
 * 编辑 message request 请求体信息
 *
 * @param id          主键ID
 * @param baseSetId   基础设置ID
 * @param reqName     request名称
 * @param reqValue    request字段值
 *
 * @return error | nil
 */
func(req *Request) EditMsgRequest(id int32, baseSetId int32, serMethodId int32, reqName string, reqValue string) error {
	reqs := &Request{
		BaseSetId: baseSetId,
		ReqName:   reqName,
		ReqValue:  reqValue,
	}

	result := reqDB().Where(&Request{Id:id}).Update(reqs)
	if result.Error != nil {
		return result.Error
	}

	//更新服务方法请求参数
	if serMethodId > 0 {
		serM := new(SerMethod).EditRerMethodById(serMethodId, reqName, "")
		if serM != nil {
			return serM
		}
	}

	return nil
}

/**
 * 根据主键ID获取一条 message request 请求体信息
 *
 * @param id  主键ID
 * @return error | nil
 */
func(req *Request) GetOneMsgRequestById(id int32) (*Request, error) {
	reqs := &Request{Id:id}
	result := reqDB().First(reqs)

	if result != nil {
		return nil, result.Error
	}

	return reqs, nil
}

/**
 * 根据基础设置ID获取 message request 请求体信息
 *
 * @param id  主键ID
 * @return error | nil
 */
func(req *Request) GetMsgRequestsByBaseId(baseSetId int32) ([]*Request, error) {
	var reqs []*Request

	result := reqDB().Find(reqs, &Request{BaseSetId:baseSetId})
	if result.Error !=nil {
		return nil, result.Error
	}

	return reqs, nil
}


func reqDB() *gorm.DB {
	return db.Database().Table("pt_msg_request")
}