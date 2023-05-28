package Controller

import (
	"StudyDemo/LoginRegisterCheck/dto"
	"StudyDemo/LoginRegisterCheck/function"
	"StudyDemo/LoginRegisterCheck/modules"
	"StudyDemo/LoginRegisterCheck/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user") //用户是通过认证的，我们能直接通过上下文获取用户的信息
	//然后将这个用户返回
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"user": dto.ToUserDto(user.(*modules.User)),
		},
	})
}

func UserRegister(ctx *gin.Context) {
	name := ctx.PostForm("name")
	user, _ := function.CheckUserByName(name)
	if user.ID > 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "用户已注册，请登录",
		})
		return
	}
	password := ctx.PostForm("password")
	//密码是不能明文保存的，创建用户的时候我们进行加密
	hasedPassword, hasedErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hasedErr != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := &modules.User{
		Name:     name,
		Password: string(hasedPassword),
	}
	err := function.SaveUser(newUser.Name, newUser.Password)
	if err != nil {
		fmt.Println("这里出现了错误")
		log.Println(err)
	}
	response.Success(ctx, nil, "注册成功")
}

func UserLogin(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	user, _ := function.CheckUserByName(name)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "该用户不存在，请先注册")
		return
	}
	//解密判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误，请输入正确的密码")
		return
	}
	//发放token
	token, err := modules.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error:%v", err)
		return
	}
	//返回结果
	//第一部分使用token使用的加密协议，第二部分储存claims中的信息，最后一部分是前面两部分加上key哈希的一个值
	response.Success(ctx, gin.H{"token": token}, "登陆成功")
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjUsImV4cCI6MTY3MjQwMTE0NywiaWF0IjoxNjcxNzk2MzQ3LCJpc3MiOiJsb2dpbklzc3VlIiwic3ViIjoidXNlciB0b2tlbiJ9.uWouR6YAbN-to5jxWmyeYW-Q1Vv7STqswhnLq5vEfOg
}
