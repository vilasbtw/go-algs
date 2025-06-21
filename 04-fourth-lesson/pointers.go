package main

import (
	"fmt"
	"go-algs/04-fourth-lesson/dog"
)

func main() {

	cerberus := dog.Dog{
		Name: "Cerberus",
	}

	cerberus.Bark(cerberus)
	cerberus.ChangeNameByReference(cerberus, "Referilson")
	cerberus.Bark(cerberus)

	cerberus.ChangeNameByValue(&cerberus, "Valorilson")
	cerberus.Bark(cerberus)

	a := 1

	// "a" memory address
	fmt.Println("'a' memory address:", &a)

	// pointer points to "a" memory address
	pointer := &a
	// prints the address of "a"
	fmt.Println("pointer memory address:", pointer)
	// prints the value stored in the memory
	fmt.Println("pointer value:", *pointer)

	a = 10
	fmt.Println("pointer value after changing 'a' value:", *pointer)

	*pointer = 100
	fmt.Println("'a' value after changing pointer value:", a)

	// change the value based on the memory address
	// all variables pointing to this address will have its value changed
	changeMemoryValue(&a)

	fmt.Printf("final pointer and 'a' value: %d, %d", *pointer, a)
}

func changeMemoryValue(a *int) {
	*a = 1000
}
