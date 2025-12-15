package handler

import (
	"net/http"
	"mitton/gin-handson/service"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetUsers())
}

func CreateUser(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	user := service.CreateUser(req.Name)
	c.JSON(http.StatusOK, user)
}
