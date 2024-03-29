package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kev87ian/BasicGolangCRUD/controllers"
	"github.com/kev87ian/BasicGolangCRUD/initializers"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:id", controllers.GetOnePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
