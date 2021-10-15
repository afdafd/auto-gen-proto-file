package model

import (
	db "customPro/protoGen/database"
	"github.com/jinzhu/gorm"
	"time"
)

type ProSet struct {
	Id            int32       `form:"id"         json:"id"`
	ProName       string      `form:"pro_name"   json:"pro_name"`
	ProPath       string      `form:"pro_path"   json:"pro_path"`
	HostName      string      `form:"host_name"  json:"host_name"`
	UserName      string      `form:"user_name"  json:"user_name"`
	Pwd           string      `form:"pwd"        json:"pwd"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	BaseSets      []*BaseSets `json:"base_sets"`
}

/**
 * 添加proto项目配置
 *
 * @param  proName  项目名称
 * @param  proPath  项目路径
 * @param  hostName 主机名
 * @param  userName 用户名
 * @param  pwd      密码
 *
 * @return nil | error
 */
func(s *ProSet) AddProSet(proName string, proPath string, hostName string, userName string, pwd string, id int32) error  {
	set := &ProSet{
		ProName:   proName,
		ProPath:   proPath,
		HostName:  hostName,
		UserName:  userName,
		Pwd:       pwd,
	}

	if result := proSetDB().Create(set); result.Error != nil {
		return result.Error
	}

	return nil
}

/**
 * 编辑proto项目配置
 *
 * @param  id       主键ID
 * @param  proName  项目名称
 * @param  proPath  项目路径
 * @param  hostName 主机名
 * @param  userName 用户名
 * @param  pwd      密码
 *
 * @return nil | error
 */
func(s *ProSet) EditProSet(id int32, proName string, proPath string, hostName string, userName string, pwd string) error {
	set := &ProSet{
		ProName:   proName,
		ProPath:   proPath,
		HostName:  hostName,
		UserName:  userName,
		Pwd:       pwd,
	}

	if result := proSetDB().Where("id = ?", id).Update(set); result.Error != nil {
		return result.Error
	}

	return nil
}

/**
 * 获取proto项目配置
 *
 * @param  id   主键ID
 * @return
 */
func(s *ProSet) GetProSetById(id int32) (*ProSet, error) {
	set:= &ProSet{Id:id}

	if result := proSetDB().First(set); result.Error != nil {
		return nil, result.Error
	}

	return set, nil
}

/**
 * 获取全部proto项目配置
 *
 * @return
 */
func(s *ProSet) GetProSets() ([]*ProSet, error) {
	var sets []*ProSet

	if result := proSetDB().Find(&sets); result.Error != nil {
		return nil, result.Error
	}

	var baseSet BaseSets
	var req Service

	for _, v := range sets {
		baseSets, err := baseSet.GetBaseSetListByProId(v.Id)
		if err != nil {
			return nil, err
		}

		v.BaseSets = baseSets

		for _, vv := range v.BaseSets {
			sers, err := req.GetServicesByBaseId(vv.Id)
			if err != nil {
				return nil, err
			}

			vv.ProtoService = sers
		}
	}

	return sets, nil
}


func proSetDB() *gorm.DB {
	return db.Database().Table("pt_pro_set")
}
