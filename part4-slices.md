# SLICES IN GO


### What is a slice?

A slice is a data structure in Go. It is used to work with sequences of data. It is similar to an array with certain, profound advantages and disadvantages:
#### The Advantages:
1) Slices are dynamically sized and, therefore, offer more flexibility when compared to arrays
2) There are various, practical built-in functions in Go that make working with slices efficient
3) They are more suitable to be used as function arguments when compared to arrays

#### The Disadvantages:
1) Slices having a dynamic size can be problem for situations in which you specifically need a data structure with a fixed size
2) They directly modify the array on which they are based. This can be a problem if you aren't careful


### Slice Declaration
There are several ways to declare slices:
1) Declaring slices through array:
```Go
var arrayOne = [4]int{14, 22, 43, 24}
sliceOne := arrayOne[0:3] // Slice from the first three elements
```
2) Declaring slices through the _make_ function:
```Go
sliceTwo := make([]float64, 4) // A slice of 4 floats, initial values are set to 0
```

3) Declaring slices like arrays (without fixed sizes)
```Go
sliceThree := []int{12, 25, 37} // Note that we did not specify a size
```

### Appending Elements to Slices
You can append elements to slices as follows:
```Go
sliceFour := []string{"Thomas", "Michael", "Jeremy"}

fmt.Println(sliceFour) // Before appending

sliceFour = append(sliceFour, "John", "William")

fmt.Println(sliceFour) // After appending
```

### Slice length
You can check the length of your slices with the _len_ function:
```Go
fmt.Println(len(sliceFour))
```

### Slice Capacity
As we have seen with slice declarations, slices can be declared based on arrays. The capacity of a slice returns the size of the original array. This can be done through the _cap_ function:
```Go
var originArray = [7]int{16, 18, 43, 54, 68, 97, 61}
slicedOrigin := originArray[1:5]
fmt.Println("The length of the slice:", len(slicedOrigin))
fmt.Println("The capacity of the slice:", cap(slicedOrigin))
```

### Array Modification - Slice Internals
Changes to a slice will directly affect the original array it depends on:
```Go
var myArr = [4]int{1, 2, 3, 4}
fmt.Println("The array before we modified the slice:", myArr)
mySlice := myArr[0:3]
mySlice[0] = 43
fmt.Println("The array after we modified the slice:", myArr)
```









   
