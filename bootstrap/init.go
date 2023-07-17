package bootstrap

import (
	"eden/conf"
	"eden/logger"
	"eden/middleware"
	"eden/router"
	"github.com/gin-gonic/gin"
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
	r := gin.Default()
	r.Use(middleware.Trace())
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())

	// router
	router.InitRouter(r)
	// bootstrap run
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	} else {
		panic(r.Run(":8080"))
	}
}
