# Data Types

One of the most important concepts of GO are the data types. You have to know how to work with them to write efficient programs. 

## Basic Data Types

### Boolean
Booleans have two possible values: true and false.

```Go
var isForSale bool = true
var isInStock = false
```

### Integers

There are many integer types (like uint8, int32 etc.). This topic will be explained further in a later part. For now, we will use only int32 and int64. The main difference between them is that int64 has more memory space and thus can be used to hold larger values.


```Go
var carNumber int32 = 976545655
var carLarger int64 = 9650846909004596
```

### Floats
There are two main float types in Go: float32 and float64. The difference between them is similar to the difference between int32 and int64.

```Go

var carMPG float32 = 43.72
var carMiles float64 = 646465.2342
```

### Strings
Strings are used for text.

```Go

var carName string = "Corolla"

```

