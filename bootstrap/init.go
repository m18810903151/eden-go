package bootstrap

import (
	"eden/conf"
	"eden/logger"
	"eden/router"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func Init() {
	// conf
	conf.InitConfig()
	// logger
	logger.Init()
	logger.SetDebugMode(true)
	// db
	db := conf.InitDb()
	// 延迟关闭
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			logger.Error("close datasource error", err.Error())
		}
	}(db)
	// router
	r := router.InitRouter()
	// bootstrap run
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	} else {
		r.Run(":8080")
	}
}
