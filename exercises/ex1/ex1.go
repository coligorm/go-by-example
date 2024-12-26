/*
Exercise 1
Create two variables a and b, and initially set them each to a different number. Write a program that swaps both values.
*/
package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Printf("a = %v, b = %v\n", a, b)
	fmt.Println("Swapping:")

	temp := a
	a = b
	b = temp

	fmt.Printf("a = %v, b = %v", a, b)
}
