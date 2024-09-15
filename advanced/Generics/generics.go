package main

import "fmt"

// A generic function to swap two values of any type
func Swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {
	// Swap two integers
	x, y := 10, 20
	x, y = Swap(x, y)
	fmt.Println("Swapped integers:", x, y)

	// Swap two strings
	str1, str2 := "Hello", "World"
	str1, str2 = Swap(str1, str2)
	fmt.Println("Swapped strings:", str1, str2)
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(v T) {
	s.elements = append(s.elements, v)
}

func (s *Stack[T]) Pop() T {
	if len(s.elements) == 0 {
		var zero T
		return zero
	}
	index := len(s.elements) - 1
	elem := s.elements[index]
	s.elements = s.elements[:index]
	return elem
}

// In the provided code, we have a function called Swap that demonstrates the use of generics in Go. Generics allow us to write functions and data structures that can work with different types without sacrificing type safety.

// The function signature func Swap[T any](a, b T) (T, T) indicates that Swap is a generic function. The [T any] syntax specifies that T can be any type. This means that Swap can be used to swap values of any type, whether it's integers, strings, or even custom types.

// Inside the function body, we simply return the values b and a, effectively swapping their positions. The return type (T, T) indicates that the function returns two values of the same type T.

// In the main function, we demonstrate the usage of the Swap function with both integers and strings. We first declare two integer variables x and y with values 10 and 20 respectively. Then, we call Swap(x, y) and assign the returned values back to x and y. This effectively swaps the values of x and y. We then print the swapped integers using fmt.Println.

// Similarly, we declare two string variables str1 and str2 with values "Hello" and "World" respectively. We call Swap(str1, str2) and assign the returned values back to str1 and str2, swapping their values. Finally, we print the swapped strings.

// The beauty of generics is that we can reuse the same Swap function for different types without having to write separate swap functions for each type. This promotes code reusability and reduces code duplication.
