package main

import (
	"fmt"
)

func main() {
	// Declaring Slices through arrays

	var arrayOne = [4]int{14, 22, 43, 24}
	sliceOne := arrayOne[0:3] // Slice from the first three elements

	fmt.Println(sliceOne)

	// Declaring slices through the make function
	sliceTwo := make([]float64, 4) // A slice of 4 floats, initial values are set to 0

	fmt.Println(sliceTwo)

	// Declaring slices like arrays (without fixed sizes)
	sliceThree := []int{12, 25, 37} // Note that we did not specify a size

	fmt.Println(sliceThree)

	// Appending elements to slices
	sliceFour := []string{"Thomas", "Michael", "Jeremy"}

	fmt.Println(sliceFour) // Before appending

	sliceFour = append(sliceFour, "John", "William")

	fmt.Println(sliceFour) // After appending

	// Slice length
	fmt.Println(len(sliceFour))

	// Slice capacity (the fixed size of the original array)
	var originArray = [7]int{16, 18, 43, 54, 68, 97, 61}
	slicedOrigin := originArray[1:5]
	fmt.Println("The length of the slice:", len(slicedOrigin))
	fmt.Println("The capacity of the slice:", cap(slicedOrigin))

	// Array modification - Slice internals
	var myArr = [4]int{1, 2, 3, 4}
	fmt.Println("The array before we modified the slice:", myArr)
	mySlice := myArr[0:3]
	mySlice[0] = 43
	fmt.Println("The array after we modified the slice:", myArr)

}
