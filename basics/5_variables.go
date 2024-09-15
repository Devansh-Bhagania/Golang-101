/*
int- stores integers (whole numbers), such as 123 or -123
float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
string - stores text, such as "Hello World". String values are surrounded by double quotes
bool- stores values with two states: true or false
*/

package main

import (
	"fmt"
)

func main() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	var a, b, c, d int = 1, 3, 5, 7 //multiple variables of the same type
	var e, f, g = 2, 4.5, "hi"      //multiple variables of different types

	//short hand declaration
	h, i, j := 1, 2, 3

	fmt.Println(a, b, c, d)
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}
