package main

import "fmt"

func main() {

	// ARRAY DECLARATION AND INITIALIZATION

	// name, [number of elements], element data type
	var arrOne [4]int     // this declares an array with 4 integer elements, all elements are 0 because we didn't initialize
	var arrText [3]string // this declares an array with 3 string elements, all elements are empty strings because we didn't initialize

	// Array initialization with values
	var arrTwo = [7]int{16, 18, 43, 54, 68, 97, 61} // an array of 7 integers (size = 7), initialized with values
	var arrTextTwo = [2]string{"Tom", "William"}    // an array of 2 strings (size = 7), initialized with values
	var arrThree = [...]int{2, 3, 4}                // an array of integers, size is not declared at first but based on the number of elements

	fmt.Println(arrOne)
	fmt.Println(arrTwo)
	fmt.Println(arrText)
	fmt.Println(arrTextTwo)
	fmt.Println(arrThree)

	// ARRAY INDICES
	fmt.Println(arrTwo[0])   // This will print out the first element of arrTwo
	fmt.Println(arrThree[2]) // This will print out the third element of arrThree

	// Changing elements through indices
	var myArr = [3]string{"John", "James", "Henry"}
	myArr[0] = "Michael" // Change the first element to "Michael"
	fmt.Println(myArr)

}
