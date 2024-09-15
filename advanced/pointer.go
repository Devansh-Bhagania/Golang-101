// Go pointers are a powerful feature that allows you to work with memory addresses directly. They are used to store the memory address of a variable. This can be useful when you need to pass a large amount of data to a function or when you need to modify the value of a variable inside a function.

package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
