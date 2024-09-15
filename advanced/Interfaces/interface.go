package main

import (
	"fmt"	package main
	
	import "fmt"
	
	// Define an interface
	type geometry interface {
		area() float64
		perim() float64
	}
	
	// Define a struct for rectangle
	type rect struct {
		width, height float64
	}
	
	// Implement the geometry interface for rect
	func (r rect) area() float64 {
		return r.width * r.height
	}
	
	func (r rect) perim() float64 {
		return 2 * (r.width + r.height)
	}
	
	// Define a struct for circle
	type circle struct {
		radius float64
	}
	
	// Implement the geometry interface for circle
	func (c circle) area() float64 {
		return 3.14 * c.radius * c.radius
	}
	
	func (c circle) perim() float64 {
		return 2 * 3.14 * c.radius
	}
	
	// Define a function that takes a geometry interface
	func measure(g geometry) {
		fmt.Println(g)
		fmt.Println(g.area())
		fmt.Println(g.perim())
	}
	
	func main() {
		r := rect{width: 3, height: 4}
		c := circle{radius: 5}
	
		measure(r)
		measure(c)
	}
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}


/*
What is an Interface?

An interface in Go is a type that specifies a set of method signatures. A type implements an interface by implementing its methods. Interfaces are a way to define behavior in Go.
Why Use Interfaces?

Abstraction: Interfaces allow you to define methods that must be implemented by any type that satisfies the interface, providing a way to abstract and generalize functionality.
Polymorphism: Interfaces enable polymorphism, allowing you to write functions that can operate on different types that implement the same interface.
Decoupling: Interfaces help decouple code by separating the definition of behavior from the implementation, making the code more modular and easier to maintain.
*/