package controllers

import (
	"firebackend/common"
	"firebackend/model"
	"github.com/gin-gonic/gin"
)

func UserInfo(cox *gin.Context) {
	var login model.LoginUser
	db := common.MysqlInit()
	db.Find(&login)
	cox.JSON(200, gin.H{
		"code":    20000,
		"msg":     "success",
		"message": "获得info成功",
		"data":    login,
	})
	defer db.Close()
}
