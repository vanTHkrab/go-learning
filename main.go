package main

import (
	"fmt"
	"go-learning/array"
	"go-learning/loops"
	"go-learning/struct_basic"
	"go-learning/variables"
)

func main() {
	fmt.Println("Go Learning - Basic Examples")

	// Variables
	fmt.Println("Variables:")
	variables.DeclareMultipleVariables()
	variables.DeclareMultipleVariablesWithInference()
	variables.MultiWordVariableNamingRules()

	// Arrays
	fmt.Println("\nArrays:")
	array.SimpleArray()
	array.MultiDimensionalArray()
	array.ArrayWithLength()

	// Structs
	fmt.Println("\nStructs:")
	struct_basic.StructBasicExample()
	struct_basic.StructWithNestedTypesExample()

	// Loops
	fmt.Println("\nLoops:")
	loops.LoopExample()
	loops.NestedLoopsExample()
	loops.LoopWithContinueAndBreak()
}
