package main

import (
	"restgo/controllers"
	"restgo/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	
	r.GET("/", controllers.ReviewsCreate)

	r.Run() // listen and serve on 0.0.0.0:8080
}