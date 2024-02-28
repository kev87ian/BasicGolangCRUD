package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kev87ian/BasicGolangCRUD/initializers"
	"github.com/kev87ian/BasicGolangCRUD/models"
	"net/http"
)

/*func PostsCreate(c *gin.Context) {
	// get data off request body
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post Created successfully"})
}*/

func PostsCreate(c *gin.Context) {

	post := models.Post{Title: "Trying it out with pointers", Body: "Hehe, a true tets to my understanding of golang"}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Post:": post})

}
