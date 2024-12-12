# VARIABLES

### Variables are declared as follows:

```go

// var keyword, variable name, variable type
var carBrand string

```


__This variable is not initialized (we did not assign a value to it with the = operator), its value is an empty string""__

```go
var carModel int

```  
 __This variable is also not initialized its value is 0__

 ```go
 
 carModel = 1970  // You can change it later on
```

__It is possible to declare and initialize variables at the same time. Note that values of variable are changeable after they are declared. This why they are called variables.__
```go

var carCity string = "New York" // We have declared a variable carCity and initialized it with the value "New York"
carCity = "Boston"              // We can change it to something else later on with the = operator
```

__It is possible to declare multiple variables at once like this:__
```go

var carMotor, carMPG float32 = 200.45, 32.7
```

 __It is also possible to declare variables like this, it is especially useful when you have several variables to declare:__

 ```go

var (
	carNameOne    string = "Corolla"
	carNameTwo    string = "Grandland"
	carNameThree  string = "Golf"
	carModelOne   int    = 2008
	carModelTwo   int    = 2022
	carModelThree int    = 2018
	)
```
__You can quickly declare and initialize variables. Note that this works only inside functions (remember that the _main_ function is also a function)__

```Go
// Short Variable Declaration with Initialization
	newCar := "Ford Fiesta" // The type is inferred from the initialized value
	fmt.Println(newCar)
```

## CONSTANTS
__Unlike variables, values of constants cannot be changed after declaration. Constants are declared as follows:__
 ```go

const carGallery string = "Gallery York"
// Trying to change it like carGallery = "Gallery Miami" would raise an error

```

__Constants can also be declared in groups:__

```go

const (
galleryYear int = 2014
galleryCity = "New York"
)


```




