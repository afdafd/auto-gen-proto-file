package common

import (
	db "customPro/protoGen/database"
	"github.com/jinzhu/gorm"
)

func GetDB(dbName string) *gorm.DB {
	return db.Database().Table(dbName)
}
