package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// init is called when the package is imported
func init() {
	// set random seed
	rand.Seed(time.Now().UnixNano())
}

// Hello is exported because it is capitalized
func Hello(name string) (string, error) {
	// return error if name is empty
	if name == "" {
		return "", errors.New("empty name")
	}

	// return greeting
	format := randomFormat()
	message := fmt.Sprintf(format, name)
	return message, nil
}

func ManyHellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// randomFormat can be used before it is declared
func randomFormat() string {
	// create slice (dynamicly sized array) of format strings
	formats := []string{
		"Hi %v!",
		"Hello %v!",
		"Welcome %v!", // every line needs a comma
	}

	// return a random greeting
	return formats[rand.Intn(len(formats))]
}
