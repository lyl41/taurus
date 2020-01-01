package dao

import (
	"fmt"
	"taurus/src/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", config.GetMysqlCfg())
	if err != nil {
		panic("open mysql err : " + err.Error())
	} else {
		fmt.Println("mysql member database Open success,", db)
	}
	db.Debug()
}

func GetDB() *gorm.DB {
	return db
}
