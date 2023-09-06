package controllers

import (
	"restgo/initializers"
	"restgo/models"

	"github.com/gin-gonic/gin"
)

func ReviewsCreate(c *gin.Context) {
	//Get the body of the request

	//Create a new review
	review := models.Review{Title: "test", Text: "test", Rating: 5, ReviewerName: "test", ReceiverID: 1}

	result := initializers.DB.Create(&review)

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
}