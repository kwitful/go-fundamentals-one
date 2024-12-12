# METHODS IN GO

### What is a method?

A method in Go is a function that is associated with a specific type. Methods allow you to define behaviors or actions for a type, making your code more organized and modular.

#### The Advantages:
1) Methods allow encapsulating functionality within a type.
2) They improve readability and maintainability of your code.
3) Methods make it easier to model real-world entities and their behaviors.

#### The Disadvantages:
1) Overusing methods can lead to overly complex and tightly coupled code.
2) Methods with poor design can reduce code clarity.

---

### Defining a Method

Methods are defined by associating a function with a specific type. Here is an example:

```Go
package main

import "fmt"

// Defining a struct
type Car struct {
	brand    string
	maxSpeed int
}

// Defining a method for Car that returns the maximum distance a car can travel in a certain time
func (carX Car) maxDistance(hours int) int {
	return carX.maxSpeed * hours
}

func main() {
	carOne := Car{brand: "BMW", maxSpeed: 160}

	// Calling the method
	fmt.Println("In two hours, the car can travel up to", carOne.maxDistance(2), "kilometers")
}
```

#### Explanation:
1) `Car` is a struct that represents a car with `brand` and `maxSpeed` fields.
2) The `maxDistance` method calculates the maximum distance the car can travel in a given number of hours.
3) The method uses a **value receiver**, which means it works on a copy of the `Car` instance.

---

### Value Receivers

A **value receiver** means the method operates on a copy of the instance it is called on. Any modifications made within the method will not affect the original instance.

#### Example:
```Go
func (carX Car) maxDistance(hours int) int {
	return carX.maxSpeed * hours
}
```
1) In this example, `carX` is a copy of the `Car` instance.
2) The original instance remains unchanged, even if `carX` is modified inside the method.

#### When to Use:
- Use value receivers when the method does not need to modify the original instance.
- Ideal for lightweight types.

---

