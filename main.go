package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sringtho/go_crud/controllers"
	"github.com/sringtho/go_crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.GetAllPosts)
	
	router.Run()
}