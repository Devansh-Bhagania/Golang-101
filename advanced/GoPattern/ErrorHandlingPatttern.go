// Adds context to errors and inspects them.

package main

import (
	"errors"
	"fmt"
	"os"
)

func readFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("readFile: %w", err)
	}
	return nil
}

func main() {
	err := readFile("nonexistent.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File does not exist")
		} else {
			fmt.Println("Error:", err)
		}
	}
}
