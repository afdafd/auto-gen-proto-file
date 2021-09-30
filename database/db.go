package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)


var db *gorm.DB

const hostName  = "127.0.0.1"
const port      = 3306
const username  = "root"
const password  = "root"
const dbName    = "ProtoFile"


func init() {
	var err error
	db, err = gorm.Open("mysql", getConnectUrl())

	if err != nil {
		log.Fatal(err.Error())
	}

	if db.Error != nil {
		fmt.Println("连接数据库失败～", db.Error)
	} else {
		fmt.Println("连接数据库成功！")
	}

	//defer db.Close()
}


func getConnectUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		hostName,
		port,
		dbName,
	)
}

func Database() *gorm.DB {
	return db
}