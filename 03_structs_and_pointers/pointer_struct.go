package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	person1 := Person{
		Name: "Doe",
	}

	// copy by value
	person2 := person1

	person2.Name = "Meyer"
	fmt.Println(person1.Name) // Doe

	// copy by reference
	person3 := &person1

	person3.Name = "Meyer"
	fmt.Println(person1.Name) // Meyer
}
