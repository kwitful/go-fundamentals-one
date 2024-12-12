// Method with a value receiver
package main

import "fmt"

// Defining a struct
type Car struct {
	brand    string
	maxSpeed int
}

// Defining a method for Car that returns the Max distance a car can travel in a certain time
func (carX Car) maxDistance(hours int) int {
	return carX.maxSpeed * hours
}

func main() {
	carOne := Car{brand: "BMW", maxSpeed: 160}

	// Calling the method
	fmt.Println("In two hours, the car can travel up to", carOne.maxDistance(2), "kilometers")
}
