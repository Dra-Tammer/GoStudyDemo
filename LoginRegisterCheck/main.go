package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	InitConfig()
	r := gin.Default()
	r = CollectRoute(r)
	//如果配置文件进行了配置，按配置进行开端口，没配置就是默认
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	//这个就是默认
	panic(r.Run())
}

func InitConfig() {
	//获取当前的工作目录
	workDir, _ := os.Getwd()
	//设置我们要读取的文件名
	viper.SetConfigName("application")
	//设置我们要读取的文件的类型
	viper.SetConfigType("yml")
	//设置文件的路径
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
}
