package config

import (
	"bytes"
	"text/template"
)

type Config struct {
	AmqpUser     string `json:"amqpUser"`
	AmqpPassword string `json:"amqpPassword"`
	AmqpHost     string `json:"amqpHost"`
	AmqpPort     string `json:"amqpPort"`
}

const amqpUrl string = "amqp://{{.AmqpUser}}:{{.AmqpPassword}}@{{.AmqpHost}}:{{.AmqpPort}}"

func (config *Config) GetAmqpConfig() string {

	var amqpConfig bytes.Buffer

	amqpTemplate, err := template.New("amqpTemplate").Parse(amqpUrl)

	if err != nil {
		panic(err)
	}

	if err := amqpTemplate.Execute(&amqpConfig, config); err != nil {
		panic(err)
	}

	return amqpConfig.String()
}
