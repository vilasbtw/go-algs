package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// tags json to lowercase all characters
type Person struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Register int    `json:"register"`
}

type Employee struct {
	Person
	Company string `json:"company_name"`
}

func (p Person) printPerson() {
	fmt.Println("Name: ", p.Name)
}

func main() {

	// first way for creating an object
	person := Person{
		Name:     "Kaique",
		Email:    "k@k.com",
		Register: 12345678900,
	}

	// second way for creating an object
	person2 := Person{"Maria", "m@m;com", 45612378900}

	person.printPerson()
	person2.printPerson()

	// Composition example
	employee := Employee{
		Person: Person{
			Name:     "John",
			Email:    "j@j.com",
			Register: 74185296300,
		},
		Company: "Claranet",
	}

	employee.printPerson()

	// Using json.Marshal to convert the struct to JSON format.
	employeeAsJson, err := json.Marshal(employee)

	if err != nil {
		log.Fatal(err.Error())
	}
	// Converting the bytes to string
	fmt.Println(string(employeeAsJson))

	jsonEmployee := `{"Name":"Carl","Email":"c@c.com","Register":486153297,"company_name":"Claranet"}`
	emptyEmployee := Employee{}

	// Putting the JSON into the employee
	// It is required to inform the emptyEmployee's memory address, otherwise, it'll change only in this scope.
	json.Unmarshal([]byte(jsonEmployee), &emptyEmployee)
	fmt.Println(emptyEmployee)
}
