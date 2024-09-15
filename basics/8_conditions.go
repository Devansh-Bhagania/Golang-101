
package main

import "fmt"

func main() {
    x := 10

    // if-else example
    if x > 5 {
        fmt.Println("x is greater than 5")
    } else {
        fmt.Println("x is less than or equal to 5")
    }

    // if-else if-else example
    if x > 15 {
        fmt.Println("x is greater than 15")
    } else if x > 5 {
        fmt.Println("x is greater than 5 but less than or equal to 15")
    } else {
        fmt.Println("x is less than or equal to 5")
    }

    // switch statement example
    day := "Tuesday"
    switch day {
    case "Monday":
        fmt.Println("Today is Monday")
    case "Tuesday":
        fmt.Println("Today is Tuesday")
    case "Wednesday":
        fmt.Println("Today is Wednesday")
    default:
        fmt.Println("It's another day")
    }

    // switch with multiple expressions example
    day = "Saturday"
    switch day {
    case "Saturday", "Sunday":
        fmt.Println("It's the weekend!")
    default:
        fmt.Println("It's a weekday")
    }

    // switch with no expression example
    switch {
    case x < 5:
        fmt.Println("x is less than 5")
    case x == 10:
        fmt.Println("x is equal to 10")
    default:
        fmt.Println("x is greater than 10")
    }
}
import "fmt"

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than or equal to 5")
	}
}
