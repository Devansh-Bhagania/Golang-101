// Effective error handling is crucial for building robust applications.

type MyError struct {
	Time    time.Time
	Message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.Time, e.Message)
}

func doSomething() error {
	return MyError{Time: time.Now(), Message: "Something went wrong"}
}




// Error Wrapping

package main

import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("not found")

func main() {
    err := fmt.Errorf("operation failed: %w", ErrNotFound)

    if errors.Is(err, ErrNotFound) {
        fmt.Println("Error is ErrNotFound")
    }
}





// Best Practices
// Use error wrapping to add context.
// Compare errors using errors.Is and errors.As.
// Avoid using panics for regular error handling.

