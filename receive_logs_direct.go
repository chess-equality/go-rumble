package main

import (
	"github.com/streadway/amqp"
	"go-rumble/util"
	"log"
	"os"
)

// To save logs to a file:
//  go run receive_logs_direct.go info warning error &>> logs_from_rabbit.log
//
// If you want to save only 'warning' and 'error' (and not 'info') log messages to a file:
//  go run receive_logs_direct.go warning error &>> logs_from_rabbit.log
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

	// Declare queue
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	util.FailOnError(err, "Failed to declare queue")

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [info] [warning] [error]", os.Args[0])
		os.Exit(0)
	}

	for _, s := range os.Args[1:] {

		log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, exchange, s)

		// Bind queue to exchange
		err = ch.QueueBind(
			q.Name,   // queue name
			s,        // routing key
			exchange, // exchange
			false,
			nil,
		)
		util.FailOnError(err, "Failed to bind queue")
	}

	// Register a consumer
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	util.FailOnError(err, "Failed to register a consumer")

	// Read messages until stopped

	forever := make(chan bool)

	go func() {

		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
