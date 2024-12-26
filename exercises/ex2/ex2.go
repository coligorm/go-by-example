/*
Exercise 2
Create a program that:
  - Reads the grades of these 3 classes (Grades range from 0 - 10)
  - Calculate the average of your grade
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter Geometry Grade : ")
	input1, _ := reader.ReadString('\n')
	// trim any whitespace
	input1 = strings.TrimSpace(input1)
	geometry_grade, _ := strconv.Atoi(input1)

	fmt.Printf("Enter Algebra Grade : ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	algebra_grade, _ := strconv.Atoi(input2)

	fmt.Printf("Enter Physics Grade : ")
	input3, _ := reader.ReadString('\n')
	input3 = strings.TrimSpace(input3)
	physics_grade, _ := strconv.Atoi(input3)

	total := geometry_grade + algebra_grade + physics_grade
	avg := total / 3

	fmt.Printf("Average Score = %v", avg)
}
