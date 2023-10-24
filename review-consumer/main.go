package main

import (
	"fmt"
	"review-consumer/configs"
	"review-consumer/controllers"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial(configs.EnvAmqpURI())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	q, err := ch.QueueDeclare(
		"ReviewQueue", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	controllers.ReviewConsumer(q, ch)
}
