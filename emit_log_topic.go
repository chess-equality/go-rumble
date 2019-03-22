package main

import (
	"github.com/streadway/amqp"
	"go-rumble/util"
	"log"
	"os"
	"strings"
)

// To emit a log with a routing key "kern.critical" type:
//  go run emit_log_topic.go "kern.critical" "A critical kernel error"
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
	exchange := "logs_topic"
	err = ch.ExchangeDeclare(
		exchange,           // name
		amqp.ExchangeTopic, // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)
	util.FailOnError(err, "Failed to declare an exchange")

	// Publish message to exchange

	body := topicBodyFrom(os.Args)

	err = ch.Publish(
		exchange,                // exchange
		routingKeyFrom(os.Args), // routing key
		false,                   // mandatory
		false,                   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	util.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func topicBodyFrom(args []string) string {

	var s string

	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}

	return s
}

func routingKeyFrom(args []string) string {

	var s string

	if (len(args) < 2) || os.Args[1] == "" {
		s = "anonymous.info"
	} else {
		s = os.Args[1]
	}

	return s
}
