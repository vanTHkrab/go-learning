package array

import (
	"fmt"
)

// SimpleArray demonstrates basic array operations in Go
func SimpleArray() {
	// Declare and initialize an array
	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	// Print the array
	fmt.Println("Array:", arr)

	// Access elements in the array
	fmt.Println("First element:", arr[0])
	fmt.Println("Second element:", arr[1])

	// Modify an element in the array
	arr[2] = 10
	fmt.Println("Modified Array:", arr)

	// Length of the array
	fmt.Println("Length of Array:", len(arr))

	// Iterate over the array
	for i, v := range arr {
		fmt.Printf("Element at index %d: %d\n", i, v)
	}
}

// MultiDimensionalArray demonstrates multi-dimensional arrays in Go
func MultiDimensionalArray() {
	// Declare and initialize a 2D array
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Print the 2D array
	fmt.Println("2D Array (Matrix):")
	for _, row := range matrix {
		fmt.Println(row)
	}

	// Access elements in the 2D array
	fmt.Println("Element at (1,1):", matrix[1][1]) // Accessing the element at row 1, column 1

	// Modify an element in the 2D array
	matrix[0][0] = 10
	fmt.Println("Modified Matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}
}

// ArrayWithLength demonstrates arrays with specified lengths
func ArrayWithLength() {
	// Declare and initialize an array with a specific length
	var arr [3]int = [3]int{10, 20, 30}

	// Print the array
	fmt.Println("Array with Length 3:", arr)

	// Access elements in the array
	fmt.Println("First element:", arr[0])
	fmt.Println("Second element:", arr[1])

	// Modify an element in the array
	arr[2] = 40
	fmt.Println("Modified Array:", arr)

	// Length of the array
	fmt.Println("Length of Array:", len(arr))

	// Iterate over the array
	for i, v := range arr {
		fmt.Printf("Element at index %d: %d\n", i, v)
	}
}
