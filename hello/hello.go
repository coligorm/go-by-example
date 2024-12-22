// Declare a main package
// a package is a way to group functions, and it's made up of all the files in the same directory
package main

// contains functions for formatting text
import (
	"fmt"

	"rsc.io/quote"
)

// Package quote collects pithy sayings.

// A main function executes by default when you run the main package.
func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}
