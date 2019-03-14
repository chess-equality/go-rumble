package main

import (
	"github.com/streadway/amqp"
	"go-rumble/util"
	"log"
	"os"
)

func main() {

	config, err := util.ReadConfig()
	if err != nil {
		panic(err)
	}

	log.Printf("amqp config = %s", config.GetAmqpConfig())

	// Connect
	conn, err := amqp.Dial(config.GetAmqpConfig())
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to create a channel")
	defer ch.Close()

	// Declare an exchange
	exchange := "logs_direct"
	err = ch.ExchangeDeclare(
		exchange,            // name
		amqp.ExchangeDirect, // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	util.FailOnError(err, "Failed to declare an exchange")

	// Publish message to exchange

	body := util.BodyFrom(os.Args)

	err = ch.Publish(
		exchange,                   // exchange
		util.SeverityFrom(os.Args), // routing key
		false,                      // mandatory
		false,                      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	util.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
