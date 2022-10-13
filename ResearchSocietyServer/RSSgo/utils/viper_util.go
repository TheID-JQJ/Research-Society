package utils

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// viper框架的封装
type ViperUtil struct {
}

func (ViperUtil) Read(file string, param string) map[string]interface{} {
	s := strings.Split(file, ".")
	fileName := s[0]
	fileSuffix := s[1]

	viper.SetConfigName(fileName)    // 文件名
	viper.SetConfigType(fileSuffix)  // 文件后缀
	viper.AddConfigPath("./config/") // 查找路径

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			log.Println("not find config file:" + fileName)
			return nil
		} else {
			// 配置文件被找到，但产生了另外的错误
			log.Println("failed to logging config file")
			return nil
		}
	}

	// 配置文件找到并成功解析
	return viper.GetStringMap(param)
}
