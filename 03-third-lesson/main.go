package main

import (
	"fmt"
	"go-algs/03-third-lesson/math"
)

type Dog struct {
	Name string
}

func (d Dog) bark() {
	fmt.Println(d.Name, "said woof!")
}

func main() {

	cerberus := Dog{
		Name: "Cerberus",
	}

	cerberus.bark()

	total := math.QuadradicOfTheTotal(1, 2, 3)
	fmt.Println(total())
}
