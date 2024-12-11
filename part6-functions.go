// BASIC FUNCTIONS

// Defining a simple function
package main

import "fmt"

// The "string" keyword that comes after the name is the data type of the parameter/argumnet, the other "string" keyword is the expected data type of the output
func greet(name string) string {
	return "Hello, " + name
}

func main() {
	// Calling the function
	message := greet("Alice")
	fmt.Println(message)
}


// Variadic Functions

// Variadic function example
package main

import "fmt"

func sum(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3, 4)) // Output: 10
    fmt.Println(sum(5, 10))      // Output: 15
}
