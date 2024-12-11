// Defining and using a struct
package main

import "fmt"

// Defining a struct
type Car struct {
	brand string
	model int
}

func main() {
	// Initializing a struct
	bmw := Car{brand: "BMW", model: 1994}

	// Printing struct fields
	fmt.Println("Car Brand:", bmw.brand)
	fmt.Println("Model:", bmw.model)
}
