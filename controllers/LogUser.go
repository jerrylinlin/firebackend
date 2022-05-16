package controllers

import (
	"firebackend/model"
	"github.com/gin-gonic/gin"
)

func LoginUser(cox *gin.Context) {
	var user model.LoginUser
	if err := cox.ShouldBind(&user); err != nil {
		cox.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		cox.JSON(200, gin.H{
			"username": user.UserName,
			"password": user.Password,
		})
	}

}
