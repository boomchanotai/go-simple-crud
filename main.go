package main

import (
	"net/http"

	"github.com/boomchanotai/golang-simple-crud/controllers"
	"github.com/boomchanotai/golang-simple-crud/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	post := router.Group("/posts")
	{
		post.GET("/", controllers.GetPosts)
		post.GET("/:id", controllers.GetPost)
		post.POST("/", controllers.CreatePost)
		post.PATCH("/:id", controllers.UpdatePost)
		post.DELETE("/:id", controllers.DeletePost)
	}

	router.Run()
}
