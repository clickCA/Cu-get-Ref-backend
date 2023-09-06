package main

import (
	"log"
	"os"
	"restgo/initializers"
	"restgo/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	initializers.LoadEnvVariables()
	// initializers.ConnectToDB()
}

func main() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	
	log.Println("Connected to mysql database",DB)
	r := gin.Default()
	
	r.POST("/reviews", func(c *gin.Context) {
		//Get the body of the request
		var body struct {
			Title string
			Text string
			Rating int
			ReviewerName string
			ReceiverID int
		}

		c.BindJSON(&body)
		//Create a new review
		review := models.Review{Title: body.Title, Text: body.Text, Rating: body.Rating, ReviewerName: body.ReviewerName, ReceiverID: body.ReceiverID}
	
		result := DB.Create(&review)
	
		if result.Error != nil {
			c.JSON(500, gin.H{
				"message": "Error creating review",
			})
			return
		}
		//Save the review
		c.JSON(200, gin.H{
			"review": review,
		})
	})

	r.GET("/reviews", func(c *gin.Context) {
		var reviews []models.Review
		DB.Find(&reviews)
		c.JSON(200, gin.H{
			"reviews": reviews,
		})
	})

	r.GET("/reviews/:id", func(c *gin.Context) {
		var review models.Review
		DB.First(&review, c.Param("id"))
		c.JSON(200, gin.H{
			"review": review,
		})
	})

	r.PUT("/reviews/:id", func(c *gin.Context) {
		var review models.Review
		DB.First(&review, c.Param("id"))
		var body struct {
			Title string
			Text string
			Rating int
			ReviewerName string
			ReceiverID int
		}
		c.BindJSON(&body)
		review.Title = body.Title
		review.Text = body.Text
		review.Rating = body.Rating
		review.ReviewerName = body.ReviewerName
		review.ReceiverID = body.ReceiverID
		DB.Save(&review)
		c.JSON(200, gin.H{
			"review": review,
		})
	})

	r.DELETE("/reviews/:id", func(c *gin.Context) {
		var review models.Review
		DB.Delete(&review, c.Param("id"))
		c.JSON(200, gin.H{
			"message": "Review deleted",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}