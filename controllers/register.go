package controllers

import (
	"github.com/gin-gonic/gin"
)

func UserReg(cox *gin.Context) {
	usn := cox.PostForm("username")
	if len(usn) == 0 {
		cox.JSON(200, gin.H{
			"code":    1,
			"message": "用户名不能为空",
		})
		return
	} else if len(usn) > 20 {
		cox.JSON(200, gin.H{
			"code":    1,
			"message": "用户名不能超过20个字符",
		})
		return
	}
	pwd := cox.PostForm("password")
	if len(pwd) == 0 {
		cox.JSON(200, gin.H{
			"code":    1,
			"message": "密码不能为空",
		})
		return
	} else if len(pwd) > 20 {
		cox.JSON(200, gin.H{
			"code":    1,
			"message": "密码不能超过20个字符",
		})
		return
	}
}
