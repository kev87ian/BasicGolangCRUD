package main

import (
	"github.com/kev87ian/BasicGolangCRUD/initializers"
	"github.com/kev87ian/BasicGolangCRUD/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
