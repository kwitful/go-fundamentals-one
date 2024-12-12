# STRUCTS IN GO

### What is a struct?

A struct in Go is a composite data type that groups together variables under a single name. These variables, called fields, can be of different types. Structs are a powerful way to define and manage complex data structures.

#### The Advantages:
1) Structs allow you to model real-world entities in your program.
2) They provide a way to group related data together.
3) Structs can contain other structs as fields, enabling nested and hierarchical data structures.

#### The Disadvantages:
1) Poor struct design can lead to bloated and inefficient code.
2) Modifying nested structs can become complex.

---

### Defining and Initializing a Struct

Here is how you define and use a basic struct in Go:

```Go
package main

import "fmt"

// Defining a struct
type Car struct {
	brand string
	model int
}

func main() {
	// Initializing a struct
	bmw := Car{brand: "BMW", model: 1994}

	// Printing struct fields
	fmt.Println("Car Brand:", bmw.brand)
	fmt.Println("Model:", bmw.model)

	// You can modify fields later on
	bmw.model = 1995
	fmt.Println("Updated model:", bmw.model)
}
```

#### Explanation:
1) The `Car` struct has two fields: `brand` of type `string` and `model` of type `int`.
2) A struct is initialized using the syntax `StructName{field1: value1, field2: value2}`.
3) Fields of a struct can be accessed or modified using the dot operator.

---

### Nested Structs

Structs in Go can contain other structs as fields. This allows for the creation of nested or hierarchical data structures.

#### Example:
```Go
package main

import "fmt"

// Defining a struct
type Car struct {
	brand string
	model int
}

type Factory struct {
	name    string
	product Car // Structs can contain other structs as fields
}

func main() {
	// Initializing a nested struct
	factoryOne := Factory{name: "Factory Alpha", product: Car{brand: "Mercedes", model: 1998}}

	fmt.Println("The name of the factory is:", factoryOne.name)
	fmt.Println("The main product of the factory is:", factoryOne.product.brand) // Accessing sub-fields
	fmt.Println("The model of the factory's products:", factoryOne.product.model)
}
```

#### Explanation:
1) The `Factory` struct has a field `product` of type `Car`.
2) Nested struct fields are accessed using the dot operator (e.g., `factoryOne.product.brand`).
3) Nested structs are useful for representing more complex entities.

---

