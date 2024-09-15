package main

import "fmt"

// Define a struct type called "Person"
type Person struct {
	name   string
	age    int
	height float64
}

func main() {
	// Create a new instance of the Person struct
	p := Person{
		name:   "John",
		age:    30,
		height: 1.75,
	}

	// Access and modify the fields of the struct
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Height:", p.height)

	// Update the values of the struct fields
	p.name = "Jane"
	p.age = 25
	p.height = 1.65

	fmt.Println("Updated Name:", p.name)
	fmt.Println("Updated Age:", p.age)
	fmt.Println("Updated Height:", p.height)
}
