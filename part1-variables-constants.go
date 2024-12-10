package main

import "fmt"

func main() {
	// VARIABLES

	// var keyword, variable name, variable type
	var carBrand string // This variable is not initialized (we did not assign a value to it with the = operator), its value is an empty string ""
	carBrand = "Mustang"

	var carModel int // This variable is also not initialized its value is 0
	carModel = 1970  // You can change it later on

	// It is possible to declare and initialize variables at the same time

	var carCity string = "New York" // We have declared a variable carCity and initialized it with the value "New York"
	carCity = "Boston"              // We can change it to something else later on with the = operator

	fmt.Println(carBrand, carModel, carCity)

	// It is possible to declare multiple variables at once like this
	var carMotor, carMPG float32 = 200.45, 32.7
	fmt.Println(carMotor, carMPG)

	// It is also possible to declare variables like this, it is especially useful when you have several variables to declare
	var (
		carNameOne    string = "Corolla"
		carNameTwo    string = "Grandland"
		carNameThree  string = "Golf"
		carModelOne   int    = 2008
		carModelTwo   int    = 2022
		carModelThree int    = 2018
	)

	fmt.Println(carNameOne, carNameTwo, carNameThree, carModelOne, carModelTwo, carModelThree)

	//CONSTANTS

	// Constant values cannot be changed after declared
	const carGallery string = "Gallery York" // Trying to change it like carGallery = "Gallery Miami" would raise an error

	// Constant can also be declared in groups like this
	const (
		galleryYear int = 2014
		galleryCity     = "New York"
	)
}
