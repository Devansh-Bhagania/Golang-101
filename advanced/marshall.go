package main

import (
	"encoding/json"
	"fmt"
)

// Person struct defines the data structure we will use for marshalling and unmarshalling.
// Struct tags (e.g., `json:"name"`) specify how the fields are encoded/decoded to/from JSON.
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	// ---------------------------
	//       MARSHALLING
	// ---------------------------

	// Use Case for Marshalling:
	// Marshalling is used to convert Go data structures (like structs) into a format
	// that can be easily transmitted over a network, stored in a file, or used in APIs.
	// JSON is a common data format for these purposes due to its readability and compatibility.

	// Create an instance of Person with sample data.
	person := Person{
		Name:    "Alice",
		Age:     28,
		Address: "789 Broadway",
	}

	// Convert the person struct into JSON format.
	// This process is called marshalling.
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	// Now jsonData holds the JSON representation of the person struct.
	// It can be sent over a network, saved to a file, or used in APIs.
	fmt.Println("Marshalled JSON:")
	fmt.Println(string(jsonData))

	// ---------------------------
	//       UNMARSHALLING
	// ---------------------------

	// Use Case for Unmarshalling:
	// Unmarshalling is used to convert data received in JSON format (e.g., from an API response,
	// file, or network) back into Go data structures for easy manipulation within the program.

	// Let's assume we received this JSON string from an external source.
	jsonString := `{"name":"Bob","age":35,"address":"456 Park Ave"}`

	// Create a variable of type Person to hold the unmarshalled data.
	var person2 Person

	// Convert the JSON string into the person2 struct.
	// This process is called unmarshalling.
	err = json.Unmarshal([]byte(jsonString), &person2)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}

	// Now person2 struct is populated with data from the JSON string.
	// We can use it just like any other Go struct.
	fmt.Println("\nUnmarshalled Struct:")
	fmt.Printf("Name: %s\n", person2.Name)
	fmt.Printf("Age: %d\n", person2.Age)
	fmt.Printf("Address: %s\n", person2.Address)
}
