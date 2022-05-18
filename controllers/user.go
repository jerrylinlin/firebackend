package controllers

import (
	"firebackend/common"
	"firebackend/model"
	"github.com/gin-gonic/gin"
)

func LoginUser(cox *gin.Context) {
	json := make(map[string]string)
	cox.BindJSON(&json)
	user := string(json["username"])
	pwd := string(json["password"])
	if len(user) == 0 || len(pwd) == 0 {
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
	db.Where("user_name = ?", user).First(&login)
	if login.UserName == user && login.Password == pwd {
		cox.JSON(200, gin.H{
			"code":    20000,
			"status":  "succeed",
			"message": "登录成功",
		})
		return
	} else {
		cox.JSON(200, gin.H{
			"code":    20001,
			"status":  "failed",
			"message": "用户名或密码不对请注册",
		})
		return
	}

}
