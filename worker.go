package main

import (
	"bytes"
	"github.com/streadway/amqp"
	"go-rumble/util"
	"log"
	"time"
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

	// Declare queue
	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	util.FailOnError(err, "Failed to declare queue")

	// QoS
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	util.FailOnError(err, "Failed to set QoS")

	// Register a consumer
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
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

			dotCount := bytes.Count(d.Body, []byte("."))
			log.Printf("Dot count: %d", dotCount)

			t := time.Duration(dotCount)

			time.Sleep(t * time.Second)

			log.Println("Done")

			d.Ack(true)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
