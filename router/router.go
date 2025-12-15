package router

import (
	"mitton/gin-handson/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/users", handler.GetUsers)
	r.POST("/users", handler.CreateUser)

	return r
}
