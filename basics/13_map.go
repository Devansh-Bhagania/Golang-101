package main

import "fmt"

func main() {
	// Creating a map using var
	var map1 map[string]int
	map1 = make(map[string]int)
	map1["apple"] = 1
	map1["banana"] = 2	package main
	
	import "fmt"
	
	func main() {
		// Creating a map using var
		var map1 map[string]int		package main
		
		import "fmt"
		
		func main() {
			// Creating maps using var and :=
			var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
			b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
		
			fmt.Printf("a\t%v\n", a)
			fmt.Printf("b\t%v\n", b)
		
			// Creating maps using make() function
			var c = make(map[string]string) // The map is empty now
			c["brand"] = "Ford"
			c["model"] = "Mustang"
			c["year"] = "1964" // c is no longer empty
			d := make(map[string]int)
			d["Oslo"] = 1
			d["Bergen"] = 2
			d["Trondheim"] = 3
			d["Stavanger"] = 4
		
			fmt.Printf("c\t%v\n", c)
			fmt.Printf("d\t%v\n", d)
		
			// Creating an empty map
			var e = make(map[string]string)
			var f map[string]string
		
			fmt.Println(e == nil) // false
			fmt.Println(f == nil) // true
		
			// Accessing map elements
			fmt.Printf(e["brand"])
		
			// Updating and adding map elements
			e["brand"] = "Ford"
			e["model"] = "Mustang"
			e["year"] = "1964"
			fmt.Println(e)
		
			e["year"] = "1970" // Updating an element
			e["color"] = "red" // Adding an element
			fmt.Println(e)
		
			// Removing element from map
			delete(e, "year")
			fmt.Println(e)
		
			// Checking for specific elements in a map
			val1, ok1 := a["brand"] // Checking for existing key and its value
			val2, ok2 := a["color"] // Checking for non-existing key and its value
			val3, ok3 := a["day"]   // Checking for existing key and its value
			_, ok4 := a["model"]    // Only checking for existing key and not its value
		
			fmt.Println(val1, ok1)
			fmt.Println(val2, ok2)
			fmt.Println(val3, ok3)
			fmt.Println(ok4)
		
			// Maps are references
			g := map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
			h := g
		
			fmt.Println(g)
			fmt.Println(h)
		
			h["year"] = "1970"
			fmt.Println("After change to h:")
			fmt.Println(g)
			fmt.Println(h)
		
			// Iterating over maps
			i := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
			for k, v := range i {
				fmt.Printf("%v : %v, ", k, v)
			}
			fmt.Println()
		
			// Iterating over maps in a specific order
			var j []string // defining the order
			j = append(j, "one", "two", "three", "four")
		
			for k, v := range i { // loop with no order
				fmt.Printf("%v : %v, ", k, v)
			}
			fmt.Println()
		
			for _, element := range j { // loop with the defined order
				fmt.Printf("%v : %v, ", element, i[element])
			}
		}
		map1 = make(map[string]int)
		map1["apple"] = 1
		map1["banana"] = 2
		fmt.Println("Map 1:", map1)
	
		// Creating a map using :=
		map2 := map[string]string{
			"red":   "#FF0000",
			"green": "#00FF00",
			"blue":  "#0000FF",
		}
		fmt.Println("Map 2:", map2)
	
		// Accessing elements
		appleCount := map1["apple"]
		fmt.Println("Apple count:", appleCount)
	
		// Checking if a key exists
		orangeCount, exists := map1["orange"]
		if exists {
			fmt.Println("Orange count:", orangeCount)
		} else {
			fmt.Println("Orange does not exist in map1")
		}
	
		// Deleting an element
		delete(map1, "banana")
		fmt.Println("Map 1 after deleting banana:", map1)
	
		// Iterating over a map
		fmt.Println("Iterating over map2:")
		for key, value := range map2 {
			fmt.Printf("Key: %s, Value: %s\n", key, value)
		}
	
		// Adding more elements to map2
		map2["yellow"] = "#FFFF00"
		map2["black"] = "#000000"
		fmt.Println("Map 2 after adding more elements:", map2)
	}
	fmt.Println("Map 1:", map1)

	// Creating a map using :=
	map2 := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	fmt.Println("Map 2:", map2)
}
