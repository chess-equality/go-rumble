package main

import (
	"github.com/streadway/amqp"
	"go-rumble/util"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	n := bodyFrom(os.Args)

	log.Printf(" [x] Requesting fib(%d)", n)

	res, err := fibonacciRPC(n)
	util.FailOnError(err, "Failed to handle RPC request")

	log.Printf(" [.] Got %d", res)
}

func fibonacciRPC(n int) (res int, err error) {

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
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	util.FailOnError(err, "Failed to declare queue")

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

	corrId := randomString(32)

	// Publish message to queue
	err = ch.Publish(
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(strconv.Itoa(n)),
		})
	util.FailOnError(err, "Failed to publish a message")

	// Check for correlation id in ack
	for d := range msgs {

		if corrId == d.CorrelationId {

			res, err = strconv.Atoi(string(d.Body))
			util.FailOnError(err, "Failed to convert body to integer")

			break
		}
	}

	return
}

func bodyFrom(args []string) int {

	var s string

	if (len(args) < 2) || os.Args[1] == "" {
		s = "30"
	} else {
		s = strings.Join(args[1:], " ")
	}
	n, err := strconv.Atoi(s)
	util.FailOnError(err, "Failed to convert arg to integer")

	return n
}

func randomString(l int) string {

	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}

	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
