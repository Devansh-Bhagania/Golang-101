package main

import "fmt"

// Function that takes two integers as parameters and returns their sum
func add(a, b int) int {
	return a + b
}

func main() {package main

import "fmt"

// Function that takes two integers as parameters and returns their sum
func add(a, b int) int {
    return a + b
}

// Function that takes a string as a parameter and prints a greeting
func greet(name string) {
    fmt.Println("Hello, " + name + "!")
}

// Function that returns multiple values
func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

// Recursive function to calculate factorial
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    // Calling the add function
    result := add(5, 3)
    fmt.Println("Sum:", result)

    // Calling the greet function
    greet("Alice")

    // Calling the divide function and handling multiple return values
    quotient, remainder := divide(10, 3)
    fmt.Println("Quotient:", quotient, "Remainder:", remainder)

    // Calling the recursive factorial function
    fact := factorial(5)
    fmt.Println("Factorial of 5:", fact)
}
	result := add(5, 3)
	fmt.Println("Sum:", result)
}
