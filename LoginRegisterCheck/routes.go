package main

import (
	"StudyDemo/LoginRegisterCheck/Controller"
	"StudyDemo/LoginRegisterCheck/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/user/register", Controller.UserRegister)
	r.POST("/user/login", Controller.UserLogin)
	//创建一个用户信息的路由
	r.GET("/user/info", middleware.AuthMiddleware(), Controller.UserInfo) //用我们的中间件保护我们的用户信息的接口
	return r
}
