package array

import (
	"fmt"
)

// Slices demonstrates basic slice operations in Go
func Slices() {
	// Declare and initialize a slice
	slice := []int{1, 2, 3, 4, 5}

	// Print the slice
	fmt.Println("Slice:", slice)

	// Access elements in the slice
	fmt.Println("First element:", slice[0])
	fmt.Println("Second element:", slice[1])

	// Modify an element in the slice
	slice[2] = 10
	fmt.Println("Modified Slice:", slice)

	// Length of the slice
	fmt.Println("Length of Slice:", len(slice))

	// Iterate over the slice
	for i, v := range slice {
		fmt.Printf("Element at index %d: %d\n", i, v)
	}

	// Append an element to the slice
	slice = append(slice, 6)
	fmt.Println("Slice after appending 6:", slice)

	// Create a new slice from an existing one
	newSlice := slice[1:4] // Slicing from index 1 to 3 (exclusive)
	fmt.Println("New Slice (from index 1 to 3):", newSlice)
}

// SliceWithMake demonstrates creating slices using the make function
func SliceWithMake() {
	// Create a slice using make
	slice := make([]int, 5) // Creates a slice of length 5 with zero values

	// Print the initial slice
	fmt.Println("Initial Slice:", slice)

	// Modify elements in the slice
	for i := range slice {
		slice[i] = i + 1 // Assign values 1 to 5
	}

	// Print the modified slice
	fmt.Println("Modified Slice:", slice)

	// Length and capacity of the slice
	fmt.Println("Length of Slice:", len(slice))
	fmt.Println("Capacity of Slice:", cap(slice))

	// Append elements to the slice
	slice = append(slice, 6, 7, 8)
	fmt.Println("Slice after appending 6, 7, 8:", slice)
}

// SliceWithCapacity demonstrates creating slices with specified capacity
func SliceWithCapacity() {
	// Create a slice with a specified capacity
	slice := make([]int, 0, 10) // Creates a slice with length 0 and capacity 10

	// Print the initial slice
	fmt.Println("Initial Slice:", slice)

	// Append elements to the slice
	for i := 1; i <= 5; i++ {
		slice = append(slice, i)
	}

	// Print the modified slice
	fmt.Println("Modified Slice:", slice)

	// Length and capacity of the slice
	fmt.Println("Length of Slice:", len(slice))
	fmt.Println("Capacity of Slice:", cap(slice))
}

// MultiDimensionalSlice demonstrates multi-dimensional slices in Go
func MultiDimensionalSlice() {
	// Declare and initialize a 2D slice
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Print the 2D slice
	fmt.Println("2D Slice (Matrix):")
	for _, row := range matrix {
		fmt.Println(row)
	}

	// Access elements in the 2D slice
	fmt.Println("Element at (1,1):", matrix[1][1]) // Accessing the element at row 1, column 1

	// Modify an element in the 2D slice
	matrix[0][0] = 10
	fmt.Println("Modified Matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}
}

// SliceWithLength demonstrates slices with specified lengths
func SliceWithLength() {
	// Declare and initialize a slice with a specific length
	slice := []int{10, 20, 30}

	// Print the slice
	fmt.Println("Slice with Length 3:", slice)

	// Access elements in the slice
	fmt.Println("First element:", slice[0])
	fmt.Println("Second element:", slice[1])

	// Modify an element in the slice
	slice[2] = 40
	fmt.Println("Modified Slice:", slice)

	// Length of the slice
	fmt.Println("Length of Slice:", len(slice))
}

// SliceWithMakeAndCapacity demonstrates creating slices with make and specified capacity
func SliceWithMakeAndCapacity() {
	// Create a slice using make with specified length and capacity
	slice := make([]int, 0, 5) // Creates a slice of length 0 with capacity 5

	// Print the initial slice
	fmt.Println("Initial Slice:", slice)

	// Append elements to the slice
	for i := 1; i <= 3; i++ {
		slice = append(slice, i)
	}

	// Print the modified slice
	fmt.Println("Modified Slice:", slice)

	// Length and capacity of the slice
	fmt.Println("Length of Slice:", len(slice))
	fmt.Println("Capacity of Slice:", cap(slice))
}
