package controllers

import (
	"firebackend/common"
	"firebackend/middleware"
	"firebackend/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoginUser(cox *gin.Context) {
	var feuser model.LoginUser
	err := cox.BindJSON(&feuser)
	if err != nil {
		cox.JSON(200, gin.H{
			"code":    20001,
			"message": "登录失败",
		})
		return
	}
	if len(feuser.UserName) == 0 || len(feuser.Password) == 0 {
		cox.JSON(200, gin.H{
			"code":    20001,
			"status":  "error",
			"message": "用户名或密码不能为空",
		})
		return
	}
	db := common.MysqlInit()
	var login model.LoginUser
	db.AutoMigrate(&login)
	db.Where("user_name = ?", feuser.UserName).First(&login)
	tokenstring, err := middleware.GenerateToken(login)
	if err != nil {
		fmt.Println("token生成失败")
		return
	}
	fmt.Println(tokenstring)
	token := map[string]string{
		"token":   tokenstring,
		"message": "登录成功",
	}
	if feuser.UserName == login.UserName && feuser.Password == login.Password {
		cox.JSON(200, gin.H{
			"code":   20000,
			"status": "succeed",
			"data":   token,
		})
		return
	} else {
		cox.JSON(200, gin.H{
			"code":    20001,
			"status":  "failed",
			"message": "用户名或密码错误请重新输入",
		})
		return
	}

}
