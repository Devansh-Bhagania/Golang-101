// The select statement lets a goroutine wait on multiple communication operations.

// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready. The select statement is just like switch statement, but in the select statement, case statement refers to communication, i.e. sent or receive operation on the channel.

package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
