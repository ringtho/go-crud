package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sringtho/go_crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		"message":"Smith Ringtho J",},
		)})
	
	router.Run()
}