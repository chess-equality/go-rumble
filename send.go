package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"go-rumble/config"
	"io/ioutil"
	"log"
)

func main() {

	config, err := readConfig()
	if err != nil {
		panic(err)
	}

	log.Printf("amqp config = %s", config.GetAmqpConfig())

	conn, err := amqp.Dial(config.GetAmqpConfig())

	failOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()
}

func failOnError(err error, msg string) {

	if err != nil {
		log.Fatalf("%s: %s", err, msg)
	}
}

func readConfig() (*config.Config, error) {

	var config config.Config

	all, err := ioutil.ReadFile("./resources/config.json")

	if err != nil {
		log.Fatal("Error reading configuration file")
		return nil, err
	}

	if errMarshall := json.Unmarshal(all, &config); errMarshall != nil {
		log.Fatalf("Error in unmarshalling: %s", errMarshall)
		return nil, errMarshall
	}

	return &config, nil
}
