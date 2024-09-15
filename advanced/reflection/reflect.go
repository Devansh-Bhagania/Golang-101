// Overview
// Reflection allows a program to inspect and modify its own structure and behavior at runtime.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("Type:", v.Type())
	fmt.Println("Kind:", v.Kind())
	fmt.Println("Value:", v.Float())
}

// Practical Uses
// Writing serialization/deserialization libraries.
// Implementing ORMs.
// Building validators.
// Performance Considerations
// Reflection is slower than direct code due to the dynamic nature.
// Use it judiciously in performance-critical applications.
