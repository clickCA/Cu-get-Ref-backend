package controllers

import (
	"net/http"
	"review-consumer/models"
	"review-consumer/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateReview(c *gin.Context) {
	var req models.ReviewRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	review := models.NewReview(req.Id, req.Reviewer, req.Message, req.Rating, time.Now())
	Publish(utils.ObjectTojson(review), "ReviewQueue", CHANNEL)
	c.JSON(http.StatusCreated, review)
}
