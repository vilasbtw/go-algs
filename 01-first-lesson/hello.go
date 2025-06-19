package main

import (
	"fmt"
	"go-algs/01-first-lesson/maths"
)

func main() {

	// WARNING: All exported variables and functions in Go must start with a capital letter.

	exported := maths.ExportedVariable
	fmt.Println(exported)

	firstNumber := 26
	secondNumber := 14

	result := maths.Sum(firstNumber, secondNumber)
	fmt.Printf("The type of the variable is: %T \n", result)
	fmt.Printf("The value of the sum is: %v", result)

}
