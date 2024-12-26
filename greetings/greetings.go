package greetings

import (
	"fmt"
	// import error module for error handling
	"errors"
	// rand implements pseudo-random number generators
	"math/rand"
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
	// Call randomFormat to return a random format greeting
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	// A slice is the same as an array, who's size changes dynamically
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
		"Hello there, %v",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	// Using rand.Intn, returns an int between 0:n (n being the length of the slice)
	return formats[rand.Intn(len(formats))]
}
