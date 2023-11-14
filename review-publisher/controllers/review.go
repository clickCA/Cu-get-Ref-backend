package controllers

import (
	"context"
	"net/http"
	"review-consumer/configs"
	"review-consumer/models"
	"review-consumer/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var reviewCollection *mongo.Collection = configs.GetCollection(configs.DB, "review")

func GetReviewsByCourseID(c *gin.Context) {
	courseID := c.Param("courseID") // Assuming courseID is a parameter in the URL

	filter := bson.M{"course_id": courseID}

	reviews, err := getReviewsByFilter(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func getReviewsByFilter(filter bson.M) ([]models.Review, error) {
	var reviews []models.Review

	cursor, err := reviewCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var review models.Review
		if err := cursor.Decode(&review); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return reviews, nil
}
func CreateReview(c *gin.Context) {
	var req models.ReviewRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	review := models.NewReviewReq(req.Id, c.Request.Method, req.Reviewer, req.Message, req.Rating, time.Now())
	Publish(utils.ObjectTojson(review), "ReviewQueue", CHANNEL)
	c.JSON(http.StatusCreated, review)
}

func UpdateReview(c *gin.Context) {
	var req models.ReviewRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	review := models.NewReviewReq(req.Id, c.Request.Method, req.Reviewer, req.Message, req.Rating, time.Now())
	Publish(utils.ObjectTojson(review), "ReviewQueue", CHANNEL)
	c.JSON(http.StatusCreated, review)
}

func DeleteReview(c *gin.Context) {
	var req models.ReviewRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	review := models.NewReviewReq(req.Id, c.Request.Method, req.Reviewer, req.Message, req.Rating, time.Now())
	Publish(utils.ObjectTojson(review), "ReviewQueue", CHANNEL)
	c.JSON(http.StatusAccepted, review)
}
