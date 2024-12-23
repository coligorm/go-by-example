package greetings

import "fmt"

// Hello returns a greeting for the named person.

// Similar to java, state parameter type and return type
// when creating a function
// Function names get captial letters
func Hello(name string) string {
	// Sprintf is a formatted string function which takes the variable
	// name, and inputs it into the %v, verb, position
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
