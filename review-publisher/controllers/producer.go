package controllers

import (
	"review-consumer/configs"
	"review-consumer/fallbacks"

	"github.com/streadway/amqp"
)

func GetConnection(url string) *amqp.Connection {
	conn, err := amqp.Dial(url)
	fallbacks.FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func GetChannel(connection *amqp.Connection) *amqp.Channel {
	ch, err := connection.Channel()
	fallbacks.FailOnError(err, "Failed to open a channel")
	return ch
}

func Publish(message, rountingKey string, channel *amqp.Channel) {
	err := channel.Publish(
		"",          // exchange
		rountingKey, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	fallbacks.FailOnError(err, "Failed to publish a message")
}

var CHANNEL = GetChannel(GetConnection(configs.EnvAmqpURI()))
