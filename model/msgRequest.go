package model

import (
	"customPro/protoGen/common"
	db "customPro/protoGen/database"
	"github.com/jinzhu/gorm"
	"time"
)

//protoMessage请求体
type Request struct {
	Id         int32       `form:"id" json:"id"`
	//BaseSets    *BaseSets    `gorm:"ForeignKey:id;AssociationForeignKey:BaseSetId"`
	BaseSetId  int32       `form:"base_set_id" json:"base_set_id"`
	ReqName    string      `form:"req_name" json:"req_name"`
	ReqValue   string      `form:"req_value[]" json:"req_value"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Fields     []map[string]string `gorm:"-"`
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
	re := &Request{
		BaseSetId: baseSetId,
		ReqName:   reqName,
		ReqValue:  reqValue,
	}

	result := reqDB().Where("id = ?", id).Update(re)
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
	re := &Request{Id:id}

	if result := reqDB().First(re); result.Error != nil {
		return nil, result.Error
	}

	for _, rValue := range common.PareJson(re.ReqValue) {
		re.Fields = append(re.Fields, rValue)
	}

	return re, nil
}

/**
 * 根据基础设置ID获取 message request 请求体信息
 *
 * @param id  主键ID
 * @return error | nil
 */
func(req *Request) GetMsgRequestsByBaseId(baseSetId int32) ([]*Request, error) {
	var reqs []*Request

	result := reqDB().Where("base_set_id = ?", baseSetId).Find(&reqs)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, v := range reqs {
		for _, rValue := range common.PareJson(v.ReqValue) {
			v.Fields = append(v.Fields, rValue)
		}
	}

	return reqs, nil
}

/**
 * 获取 message request 请求体信息
 *
 * @param id  主键ID
 * @return error | nil
 */
func (req *Request) GetAllMsgRequests() ([]*Request, error)  {
	var reqs []*Request

	result := reqDB().Find(&reqs)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, v := range reqs {
		for _, rValue := range common.PareJson(v.ReqValue) {
			v.Fields = append(v.Fields, rValue)
		}
	}

	return reqs, nil
}


func reqDB() *gorm.DB {
	return db.Database().Table("pt_msg_request")
}