package main

import (
	"fmt"
)

func main() {

	// If - else if - else statements

	var carModel int = 1986

	if carModel > 1972 {
		fmt.Println("The car isn't old")
	} else if carModel > 2010 {
		fmt.Println("The car is quite new")
	} else {
		fmt.Println("The car is old")
	}

	// Switch
	var carCity string = "New York"

	switch carCity {
	case "New York":
		fmt.Println("The car is in USA")

	case "London":
		fmt.Println("The car is in UK")

	case "Istanbul":
		fmt.Println("The car is in Turkiye")

	}

	// Iterating over items

	carStore := []string{"Boosterz", "Exhauster", "Greaser"}

	// With index
	for index, value := range carStore {
		fmt.Printf("The index of the store is %q and its name is %q", index, value)
	}

	// Without index
	for _, value := range carStore {
		fmt.Printf("\nThe name of the store is %q", value)
	}

}
