package variables

import (
	"fmt"
)

// DeclareVariables demonstrates variable declaration in Go
func DeclareMultipleVariables() {
	// Declare multiple variables of different types
	var a, b, c int = 1, 2, 3
	var x, y, z float64 = 1.1, 2.2, 3.3
	var isTrue, isFalse bool = true, false
	var message, greeting string = "Hello", "World"

	// Print the variables
	fmt.Println("Integers:", a, b, c)
	fmt.Println("Floats:", x, y, z)
	fmt.Println("Booleans:", isTrue, isFalse)
	fmt.Println("Strings:", message, greeting)
	// Declare constants
	const pi, e = 3.14159, 2.71828
	fmt.Println("Constants Pi and e:", pi, e)
}

// DeclareMultipleVariablesWithInference demonstrates variable declaration with type inference
func DeclareMultipleVariablesWithInference() {
	// Declare multiple variables with type inference
	a, b, c := 1, 2, 3
	x, y, z := 1.1, 2.2, 3.3
	isTrue, isFalse := true, false
	message, greeting := "Hello", "World"

	// Print the variables
	fmt.Println("Integers:", a, b, c)
	fmt.Println("Floats:", x, y, z)
	fmt.Println("Booleans:", isTrue, isFalse)
	fmt.Println("Strings:", message, greeting)
	// Declare constants with type inference
	const pi, e = 3.14159, 2.71828
	fmt.Println("Constants Pi and e:", pi, e)
}
