package route

import (
	"firebackend/controllers"
	"github.com/gin-gonic/gin"
)

func RoutRegister() *gin.Engine {
	R := gin.Default()
	err := R.SetTrustedProxies([]string{"192.168.1.2"})
	if err = nil; err != nil {
		panic(err)
	}
	//R.Use(gin.Logger())
	//R.Use(gin.Recovery())
	V := R.Group("/api/v1")
	{
		V.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		V.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		V.GET("/user/info", controllers.UserInfo)
	}
	{
		V.POST("/user/login", controllers.LoginUser)
		V.POST("/register", controllers.UserReg)
		V.POST("/logout", func(cox *gin.Context) {
			cox.JSON(200, gin.H{
				"message": "logout",
			})
		})
	}
	return R
}
