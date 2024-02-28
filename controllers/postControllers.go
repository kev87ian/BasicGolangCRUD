package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kev87ian/BasicGolangCRUD/initializers"
	"github.com/kev87ian/BasicGolangCRUD/models"
	"log"
	"net/http"
)

func PostsCreate(c *gin.Context) {
	// Define a struct to hold request body data
	var body struct {
		Body  string
		Title string
	}
	// Bind request body JSON to the 'body' struct
	err := c.Bind(&body)
	if err != nil {
		log.Fatal(err.Error()) // Log any binding errors
		return                 // Return early if there's an error
	}
	// Create a new Post instance using data from the 'body' struct
	post := models.Post{Title: body.Title, Body: body.Body}

	// Create the post in the database
	result := initializers.DB.Create(&post)
	// Check if there was an error during database operation
	if result.Error != nil {
		c.Status(400) // Set HTTP status code to 400 (Bad Request)
		return        // Return early if there's an error
	}

	// Return JSON response with the created post
	c.JSON(http.StatusCreated, gin.H{"Post:": post})
}

func GetPosts(c *gin.Context) {

}
func PostsIndex(c *gin.Context) {
	// get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}
