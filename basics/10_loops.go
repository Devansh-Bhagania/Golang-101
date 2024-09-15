package main

import "fmt"

func main() {
	// For loop
	for i := 1; i <= 5; i++ {
		fmt.Println("For loop iteration:", i)
	}

	// While loop equivalent	package main
	
	import "fmt"
	
	func main() {
		// For loop
		for i := 1; i <= 5; i++ {
			fmt.Println("For loop iteration:", i)
		}
	
		// While loop equivalent
		j := 1
		for j <= 5 {
			fmt.Println("While loop iteration:", j)
			j++
		}
	
		// Infinite loop with break statement
		k := 1
		for {
			fmt.Println("Infinite loop iteration:", k)
			k++
			if k > 5 {
				break
			}
		}
	
		// Range loop over a slice
		numbers := []int{1, 2, 3, 4, 5}
		for index, value := range numbers {
			fmt.Printf("Range loop iteration: index=%d, value=%d\n", index, value)
		}
	
		// Break and continue example
		for l := 1; l <= 10; l++ {
			if l == 5 {
				fmt.Println("Breaking at iteration:", l)
				break
			}
			if l%2 == 0 {
				fmt.Println("Continuing at iteration:", l)
				continue
			}
			fmt.Println("Loop iteration:", l)
		}
	}
	j := 1
	for j <= 5 {
		fmt.Println("While loop iteration:", j)
		j++
	}

	// Infinite loop with break statement
	k := 1
	for {
		fmt.Println("Infinite loop iteration:", k)
		k++
		if k > 5 {
			break
		}
	}

	// Looping over a collection
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
