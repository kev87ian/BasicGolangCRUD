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
	err := r.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
