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
func GetAllPosts(c *gin.Context) {
	// get the posts
	var posts []models.Post
	if result := initializers.DB.Find(&posts); result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"Posts": posts})
		return
	}
	c.JSON(200, gin.H{
		"Total_Posts": len(posts),
		"Posts":       posts,
	})
}

func GetOnePost(c *gin.Context) {
	// get id from parameter
	id := c.Param("id")

	// declare a "post" variable to hold the retrieved post
	var post models.Post
	// If the post is not found, return a JSON response with a 404 status
	if result := initializers.DB.First(&post, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Post does not exist",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}
func UpdatePost(c *gin.Context) {
	// get id off the url
	id := c.Param("id")
	// get the data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	// find the post we're updating
	var post models.Post
	initializers.DB.First(&post, id)
	//update it
	db := initializers.DB

	if result := db.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body}); result.Error != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"Error": "Post was not updated",
		})
		return
	}
	//respond with it

	c.JSON(http.StatusOK, gin.H{
		"Post": post,
	})
}
func DeletePost(c *gin.Context) {
	// Get the ID from the URL parameter
	id := c.Param("id")
	db := initializers.DB

	// Declare a variable to hold the post
	var post models.Post

	// Find the post by ID
	if result := db.First(&post, id); result.Error != nil {
		// If the post is not found, return a JSON response with a 404 status code
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Delete the post
	if result := db.Delete(&post); result.Error != nil {
		// If an error occurs during deletion, return a JSON response with a 500 status code
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
