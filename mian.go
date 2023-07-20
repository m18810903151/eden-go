package main

import (
	"eden/common"
	"eden/conf"
	"eden/middleware"
	"eden/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {

	// conf
	config := conf.InitConf()
	// db gorm
	db := common.InitDb(config.SqlConn)
	// logger
	common.InitLogger(config.LoggerConf)
	// 延迟关闭
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("close datasource error", err.Error())
		}
	}(db)
	r := gin.New()
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	r.Use(middleware.Cors())
	// routers
	routers.InitRouters(r)
	// bootstrap run
	port := config.AppConf.Port
	if port != "" {
		panic(r.Run(":" + port))
	} else {
		panic(r.Run(":8080"))
	}
}
