package main

import (
	"fmt"
	"review-consumer/controllers"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/reviews", func(ctx *gin.Context) {
		controllers.GetAllReviews(ctx)
	})

	router.GET("/reviews/:id", func(ctx *gin.Context) {
		controllers.GetReview(ctx)
	})

	router.POST("/reviews", func(ctx *gin.Context) {
		controllers.CreateReview(ctx)
	})

	router.DELETE("/reviews", func(ctx *gin.Context) {
		controllers.DeleteReview(ctx)
	})

	fmt.Println("running in :8080")
	router.Run(":8080")
}
