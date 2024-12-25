// Declare a main package
// a package is a way to group functions, and it's made up of all the files in the same directory
package main

import (
	// contains functions for formatting text
	"fmt"
	// implements a simple logging package
	"log"

	"rsc.io/quote"

	// importing my own module greeting
	"go-by-example/greetings"
)

// Package quote collects pithy sayings.

// A main function executes by default when you run the main package.
func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// passing the name "coligorm" into the Hello function
	// from the greetings module

	// I need to Edit the go-by-example/hello module to use my local go-by-example/greetings module.
	// by running go mod edit -replace go-by-example/greetings=../greetings
	name_message, err := greetings.Hello("coligorm")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name_message)

	// Request a error greeting message.
	blank_message, err := greetings.Hello("")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(blank_message)
}
