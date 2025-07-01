package loops

import (
	"fmt"
)

// LoopExample demonstrates the use of different types of loops in Go
func LoopExample() {
	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println("For loop iteration:", i)
	}

	// While-like loop (using for)
	j := 0
	for j < 5 {
		fmt.Println("While-like loop iteration:", j)
		j++
	}

	// Range loop over a slice
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Range loop - Index: %d, Value: %d\n", index, value)
	}

	// Infinite loop (use with caution)
	k := 0
	for {
		if k >= 5 {
			break // Break out of the infinite loop
		}
		fmt.Println("Infinite loop iteration:", k)
		k++
	}
}

// Nested loops example
func NestedLoopsExample() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Nested loop - i: %d, j: %d\n", i, j)
		}
	}
}

// LoopWithContinueAndBreak demonstrates the use of continue and break statements in loops
func LoopWithContinueAndBreak() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		if i >= 7 {
			break // Exit the loop if i is 7 or greater
		}
		fmt.Println("Odd number:", i)
	}
}

// LoopWithLabels demonstrates the use of labels with loops
func LoopWithLabels() {
outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("Breaking out of outer loop at i=1, j=1")
				break outerLoop // Breaks out of the outer loop
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
}

// LoopWithRangeAndMap demonstrates looping over a map using range
func LoopWithRangeAndMap() {
	m := map[string]int{
		"apple":  5,
		"banana": 3,
		"cherry": 8,
	}

	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}

// LoopWithRangeAndString demonstrates looping over a string using range
func LoopWithRangeAndString() {
	str := "Hello, World!"
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}

// LoopWithRangeAndArray demonstrates looping over an array using range
func LoopWithRangeAndArray() {
	arr := [5]int{10, 20, 30, 40, 50}
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

// LoopWithRangeAndSlice demonstrates looping over a slice using range
func LoopWithRangeAndSlice() {
	slice := []string{"Go", "is", "fun"}
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
}
