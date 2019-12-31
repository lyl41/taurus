package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var db *gorm.DB

func init() {
	var err error
	//TODO
	db, err = gorm.Open("mysql", "root:catron&*()@tcp(localhost:3306)/member?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("open mysql err : ", err)
		os.Exit(0)
	} else {
		fmt.Println("mysql member database Open success,", db)
	}
	db.Debug()
}

func GetDB() *gorm.DB {
	return db
}

