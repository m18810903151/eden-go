package main

import (
	"eden/conf"
	"eden/dao"
	"eden/logger"
	"eden/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func main() {
	conf.InitConfig()
	engine := gin.Default()
	logger.Setup()
	logger.SetDebugMode(true)
	db := conf.InitDb()

	// 延迟关闭
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			logger.Error("close datasource error", err.Error())
		}
	}(db)
	engine.GET("user/info", func(context *gin.Context) {
		util.ResponseOK(context, dao.QueryUserById(1))
	})

	port := viper.GetString("server.port")
	if port != "" {
		panic(engine.Run(":" + port))
	} else {
		engine.Run(":8080")
	}
}
