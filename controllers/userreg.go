package controllers

import "github.com/gin-gonic/gin"

func UserReg(cox *gin.Context) {
	user := cox.PostForm("username")
	if user == "" {
		cox.JSON(200, gin.H{
			"code": 20001,
			"msg":  "用户名不能为空",
		})
		return
	}
}
