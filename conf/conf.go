package conf

import (
	"github.com/spf13/viper"
	"os"
)

// InitConfig 初始化config文件
func InitConfig() {
	workDir, _ := os.Getwd()
	// 读取的文件名
	viper.SetConfigName("conf")
	// 读取的文件类型
	viper.SetConfigType("yml")
	// 读取的路径
	viper.AddConfigPath(workDir + "/conf")

	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}
