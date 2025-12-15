package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Railsでいう app/models/user.rb
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// とりあえず仮の「DB」代わり（メモリ上の配列）
var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

func main() {
	r := gin.Default()

	// health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// GET /users （index）
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	// POST /users （create）
	r.POST("/users", func(c *gin.Context) {
		var req struct {
			Name string `json:"name" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "name is required",
			})
			return
		}

		newID := users[len(users)-1].ID + 1
		user := User{
			ID:   newID,
			Name: req.Name,
		}
		users = append(users, user)

		c.JSON(http.StatusCreated, user)
	})

	r.Run(":3000")
}
