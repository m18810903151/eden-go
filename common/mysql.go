package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDb(sqlConn string) *gorm.DB {
	db, err := gorm.Open("mysql", sqlConn)
	if err != nil {
		fmt.Println("failed to init mysql")
	}
	DB = db
	return DB
}
