package main

import (
	"fmt"
	"log"
	"os"

	"example.com/greetings"
)

// main is called when you run the package
func main() {
	// configure logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0) // no timestamp

	// get command line args
	if len(os.Args) == 1 { // first arg is path
		log.Fatal("no username was provided")
		// NOTE: "" is not passed, so greetings.Hello cannot fail
	}
	usernames := os.Args[1:]

	// greet users
	messages, err := greetings.ManyHellos(usernames)
	if err != nil {
		log.Fatal(err)
	}
	for _, message := range messages {
		fmt.Println(message)
	}
}
