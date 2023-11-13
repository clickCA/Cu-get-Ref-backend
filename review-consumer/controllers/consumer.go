package controllers

import (
	"fmt"
	"review-consumer/configs"
	"review-consumer/fallbacks"
	"review-consumer/utils"
	"runtime"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

var reviewCollection *mongo.Collection = configs.GetCollection(configs.DB, "review")

func ReviewConsumer(queue amqp.Queue, channel *amqp.Channel) {

	msgs, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	fallbacks.FailOnError(err, "Failed to register a consumer")

	runtime.GOMAXPROCS(8)
	forever := make(chan bool)
	fmt.Println("Waiting for message")
	go func() {
		for d := range msgs {
			reqObj := utils.JsonToObject(([]byte(d.Body)))

			switch method := reqObj.METHOD; method {

			case "POST":
				go configs.InsertOneObj(reqObj, reviewCollection)

			case "DELETE":
				fmt.Println(reqObj)
				go configs.DeleteOneObj(reqObj, reviewCollection)

			default:
				fmt.Println("Invalid Method")
			}

		}
	}()
	<-forever

}

func CourseConsumer(queue amqp.Queue, channel *amqp.Channel) {

	msgs, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	fallbacks.FailOnError(err, "Failed to register a consumer")

	runtime.GOMAXPROCS(8)
	forever := make(chan bool)
	fmt.Println("Waiting for message")
	go func() {
		for d := range msgs {
			reqObj := utils.JsonToObject(([]byte(d.Body)))

			switch method := reqObj.METHOD; method {

			case "POST":
				go configs.InsertOneObj(reqObj, reviewCollection)

			case "DELETE":
				go configs.DeleteOneObj(reqObj, reviewCollection)

			default:
				fmt.Println("Invalid Method")
			}

		}
	}()
	<-forever

}
