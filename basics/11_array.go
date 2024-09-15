/*
var array_name = [length]datatype{values} // here length is defined

or

var array_name = [...]datatype{values} // here length is inferred

*/

// with :=

/*
array_name := [length]datatype{values} // here length is defined

or

array_name := [...]datatype{values} // here length is inferred

*/

package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 2, 3} // or [...]int{1,2,3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	prices := [3]int{10, 20, 30}

	fmt.Println(arr1)
	fmt.Println(arr2)
}
