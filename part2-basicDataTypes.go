package main

import "fmt"

func dataShow() {

	// Booleans
	var isForSale bool = true
	var isInStock bool = false

	fmt.Println(isForSale, isInStock)

	// Integers
	var carNumber int32 = 976545655
	var carLarger int64 = 9650846909004596

	fmt.Println(carNumber, carLarger)

	// Floats

	var carMPG float32 = 43.72
	var carMiles float64 = 646465.2342

	fmt.Println(carMPG, carMiles)

	// Strings

	var carName string = "Corolla"

	fmt.Println(carName)

}
