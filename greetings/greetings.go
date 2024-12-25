package greetings

import (
	"fmt"
	// import error module for error handling
	"errors"
)

// Hello returns a greeting for the named person.

// Similar to java, state parameter type and return type
// when creating a function
// Function names get captial letters
func Hello(name string) (string, error) {
	// If no name is given, return an error with a message
	if name == "" {
		// New returns an error that formats as the given text
		return "", errors.New("empty name")
	}
	// Sprintf is a formatted string function which takes the variable
	// name, and inputs it into the %v, verb, position
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message, nil
}
