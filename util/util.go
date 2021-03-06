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

	log.Printf(" (BodyFrom) body = %s", s)

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

	log.Printf(" (SeverityFrom) severity = %s", s)

	return s
}

func Append(slice, data []byte) []byte {

	l := len(slice)

	if l+len(data) > cap(slice) { // reallocate

		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data))*2)

		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)

		slice = newSlice
	}

	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)

	return slice
}
