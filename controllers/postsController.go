package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sringtho/go_crud/initializers"
	"github.com/sringtho/go_crud/models"
)

func CreatePost(c *gin.Context) {
	// Get Data off request body
	var body struct {
		Body string
		Title string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind data",
		})
	}

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create post",
		})
		return
	}

	// return
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})

}

func GetAllPosts(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func GetSinglePost(c *gin.Context) {
	// Get post id from route
	id := c.Param("id")
	// get the post from db
	var post models.Post
	initializers.DB.Find(&post, id)
	// respond with the post
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}