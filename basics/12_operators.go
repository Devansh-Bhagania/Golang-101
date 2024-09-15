package main

import "fmt"

func main() {
	// Arithmetic operators
	a := 10
	b := 5
	fmt.Println("Addition:", a+b)       // Addition: 15
	fmt.Println("Subtraction:", a-b)    // Subtraction: 5
	fmt.Println("Multiplication:", a*b) // Multiplication: 50
	fmt.Println("Division:", a/b)       // Division: 2
	fmt.Println("Modulus:", a%b)        // Modulus: 0

	// Comparison operators
	fmt.Println("Equal to:", a == b)                 // Equal to: false
	fmt.Println("Not equal to:", a != b)             // Not equal to: true
	fmt.Println("Greater than:", a > b)              // Greater than: true
	fmt.Println("Less than:", a < b)                 // Less than: false
	fmt.Println("Greater than or equal to:", a >= b) // Greater than or equal to: true
	fmt.Println("Less than or equal to:", a <= b)    // Less than or equal to: false

	// Logical operators
	x := true
	y := false
	fmt.Println("Logical AND:", x && y) // Logical AND: false
	fmt.Println("Logical OR:", x || y)  // Logical OR: true
	fmt.Println("Logical NOT:", !x)     // Logical NOT: false

	// Assignment operators
	c := 20
	c += 5                                       // c = c + 5
	fmt.Println("Addition assignment:", c)       // Addition assignment: 25
	c -= 10                                      // c = c - 10
	fmt.Println("Subtraction assignment:", c)    // Subtraction assignment: 15
	c *= 2                                       // c = c * 2
	fmt.Println("Multiplication assignment:", c) // Multiplication assignment: 30
	c /= 3                                       // c = c / 3
	fmt.Println("Division assignment:", c)       // Division assignment: 10
	c %= 4                                       // c = c % 4
	fmt.Println("Modulus assignment:", c)        // Modulus assignment: 2
}
