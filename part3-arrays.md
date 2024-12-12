# ARRAYS IN GO

### What is an Array?
An array is a data structure, with a fixed size, that is used to store data of the same type. 

### Array Declaration (without initialization)
Arrays in Go are declared as follows:

```Go
// name, [number of elements], element data type
	var arrOne [4]int     // this declares an array with 4 integer elements, all elements are 0 because we didn't initialize
	var arrText [3]string // this declares an array with 3 string elements, all elements are empty strings because we didn't initialize
```
### Array Initialization
Arrays in Go can be initialized with values as follows:
```Go
// Array initialization with values
	var arrTwo = [7]int{16, 18, 43, 54, 68, 97, 61} // an array of 7 integers (size = 7), initialized with values
	var arrTextTwo = [2]string{"Tom", "William"}    // an array of 2 strings (size = 7), initialized with values
	var arrThree = [...]int{2, 3, 4}                // an array of integers, size is not declared at first but based on the number of elements
```

### Array Indices
You can access the index of an array element. Note that Go, like many other programmin languages, starts counting index numbers from 0.

```Go
// Array Indices
	fmt.Println(arrTwo[0])   // This will print out the first element of arrTwo
	fmt.Println(arrThree[2]) // This will print out the third element of arrThree
```

### Changing Array Elements (through indices)
You can change the elements of the array by accessing them through their index number and assigning a new value, as follows:
```Go
// Changing elements through indices
	var myArr = [3]string{"John", "James", "Henry"}
	myArr[0] = "Michael" // Change the first element to "Michael"
	fmt.Println(myArr)

```
