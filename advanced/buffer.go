/*
Buffer
The buffer belongs to the byte package of the Go language, and we can use these package to manipulate the byte of the string.
*/

package main

import (
	"bytes"
	"fmt"
)

//importing the bytes package so that buffer can be used

func main() {
	//Creating buffer variable to hold and manage the string data
	var strBuffer bytes.Buffer
	fmt.Fprintf(&strBuffer, "It is number: %d, This is a string: %v\n", 10, "Bridge")
	strBuffer.WriteString("[DONE]")
	fmt.Println("The string buffer output is", strBuffer.String())
}
