package variables

import (
	"fmt"
)

// DeclareVariables demonstrates variable declaration in Go
func DeclareVariables() {
	// Declare variables of different types
	var integer int = 42
	var floatingPoint float64 = 3.14
	var boolean bool = true
	var str string = "Hello, World!"

	// Print the variables
	fmt.Println("Integer:", integer)
	fmt.Println("Floating Point:", floatingPoint)
	fmt.Println("Boolean:", boolean)
	fmt.Println("String:", str)

	// Declare a constant
	const pi = 3.14159
	fmt.Println("Constant Pi:", pi)
}

// DeclareVariablesWithInference Type is inferred from the value
func DeclareVariablesWithInference() {
	// Declare variables with type inference
	integer := 42
	floatingPoint := 3.14
	boolean := true
	str := "Hello, World!"

	// Print the variables
	fmt.Println("Integer:", integer)
	fmt.Println("Floating Point:", floatingPoint)
	fmt.Println("Boolean:", boolean)
	fmt.Println("String:", str)

	// Declare a constant with type inference
	const pi = 3.14159
	fmt.Println("Constant Pi:", pi)
}

/* Difference between var and :=
In Go, the `var` keyword is used to declare a variable with an explicit type, while the `:=` operator is used for short variable declaration where the type is inferred from the value assigned.
*/
