// Declare a main package
// a package is a way to group functions, and it's made up of all the files in the same directory
package main

// contains functions for formatting text
import (
	"fmt"

	"rsc.io/quote"

	// importing my own module greeting
	"go-by-example/greetings"
)

// Package quote collects pithy sayings.

// A main function executes by default when you run the main package.
func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	// passing the name "coligorm" into the Hello function
	// from the greetings module

	// I need to Edit the go-by-example/hello module to use my local go-by-example/greetings module.
	// by running go mod edit -replace go-by-example/greetings=../greetings
	message := greetings.Hello("coligorm")
	fmt.Println(message)
}
