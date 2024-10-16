package main

import (
	"github.com/sringtho/go_crud/initializers"
	"github.com/sringtho/go_crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}