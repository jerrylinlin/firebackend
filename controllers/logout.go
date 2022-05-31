package controllers

import (
	"github.com/gin-gonic/gin"
)

func LogOut(cox *gin.Context) {
	cox.JSON(200, gin.H{
		"code":    20000,
		"msg":     "success",
		"message": "注销成功",
	})
}
