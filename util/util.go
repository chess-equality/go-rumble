package util

import (
	"encoding/json"
	"go-rumble/config"
	"io/ioutil"
	"log"
	"strings"
)

const (
	SeverityInfo    = "info"
	SeverityWarning = "warning"
	SeverityError   = "error"
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

func BodyFrom(args []string) string {

	var s string

	if (len(args) < 2) || args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}

	log.Printf(" s = %s", s)

	return s
}

func SeverityFrom(args []string) string {

	var s string

	if (len(args) < 2) || args[1] == "" {
		s = SeverityInfo
	} else {
		switch args[1] {
		case SeverityInfo:
		default:
			s = SeverityInfo
		case SeverityWarning:
			s = SeverityWarning
		case SeverityError:
			s = SeverityError
		}
	}

	log.Printf(" s = %s", s)

	return s
}
