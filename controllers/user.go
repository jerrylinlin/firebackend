package controllers

import (
	"firebackend/common"
	"firebackend/model"
	"github.com/gin-gonic/gin"
)

func LoginUser(cox *gin.Context) {
	json := make(map[string]string)
	err := cox.BindJSON(&json)
	if err != nil {
		cox.JSON(200, gin.H{
			"code":    20001,
			"message": "登录失败",
		})
		return
	}
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
	token := map[string]string{
		"token":   "default",
		"message": "登录成功",
	}
	if login.UserName == user && login.Password == pwd {
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
