package dog

import "fmt"

type Dog struct {
	Name string
}

func (Dog) Bark(d Dog) {
	fmt.Println(d.Name, "says woof!")
}

func (Dog) ChangeNameByReference(d Dog, name string) {
	fmt.Println("Previous name: ", d.Name)

	d.Name = name
	fmt.Println("New reference name: ", d.Name)
}

// when a pointer is an argument, it changes its value, working like a global variable
func (Dog) ChangeNameByValue(d *Dog, name string) {
	d.Name = name
}
