package util

import (
	"encoding/json"
	"go-rumble/config"
	"io/ioutil"
	"log"
)

func FailOnError(err error, msg string) {

	if err != nil {
		log.Fatalf("%s: %s", err, msg)
	}
}

func ReadConfig() (*config.Config, error) {

	var config config.Config

	all, err := ioutil.ReadFile("./resources/config.json")

	if err != nil {
		log.Fatalf("Error reading configuration file: %s", err)
		return nil, err
	}

	if err := json.Unmarshal(all, &config); err != nil {
		log.Fatalf("Error in unmarshalling: %s", err)
		return nil, err
	}

	return &config, nil
}
