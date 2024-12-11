// Defining and using a struct
package main

import "fmt"

// Defining a struct
type Car struct {
	brand string
	model int
}

type Factory struct {
	name    string
	product Car // Structs can contain other structs as fields
}

func main() {
	// Initializing a struct
	bmw := Car{brand: "Mercedes", model: 1994}

	// Printing struct fields
	fmt.Println("Car Brand:", bmw.brand)
	fmt.Println("Model:", bmw.model)

	// You can modify fields later on
	bmw.model = 1995
	fmt.Println("Updated model:", bmw.model)

	// NESTED STRUCTS

	factoryOne := Factory{name: "Factory Alpha", product: Car{brand: "Mercedes", model: 1998}}

	fmt.Println("The name of the factory is:", factoryOne.name)
	fmt.Println("The main product of the factory is:", factoryOne.product.brand) // You can access sub-fields of nested structs like this
	fmt.Println("The model of the factory's products:", factoryOne.product.model)
}
