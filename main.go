package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kev87ian/BasicGolangCRUD/initializers"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello friend!",
		})
	})

	r.Run()
}
