package maps

import (
	"fmt"
)

// MapExample demonstrates basic operations on a map in Go
func MapExample() {
	// Declare and initialize a map
	myMap := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}

	// Print the map
	fmt.Println("Map:", myMap)

	// Access elements in the map
	fmt.Println("Apple count:", myMap["apple"])
	fmt.Println("Banana count:", myMap["banana"])

	// Modify an element in the map
	myMap["banana"] = 10
	fmt.Println("Modified Map:", myMap)

	// Add a new element to the map
	myMap["grape"] = 4
	fmt.Println("Map after adding grape:", myMap)

	// Delete an element from the map
	delete(myMap, "orange")
	fmt.Println("Map after deleting orange:", myMap)

	// Iterate over the map
	for key, value := range myMap {
		fmt.Printf("%s: %d\n", key, value)
	}

	// Check if a key exists in the map
	if count, exists := myMap["apple"]; exists {
		fmt.Printf("Apple exists with count: %d\n", count)
	} else {
		fmt.Println("Apple does not exist in the map.")
	}
}

// MapWithMake demonstrates creating a map using the make function
func MapWithMake() {
	// Create a map using make
	myMap := make(map[string]int)

	// Add elements to the map
	myMap["cat"] = 2
	myMap["dog"] = 3
	myMap["bird"] = 5

	// Print the map
	fmt.Println("Map created with make:", myMap)

	// Access an element in the map
	fmt.Println("Dog count:", myMap["dog"])

	// Modify an element in the map
	myMap["cat"] = 4
	fmt.Println("Modified Map:", myMap)

	// Check if a key exists in the map
	if count, exists := myMap["bird"]; exists {
		fmt.Printf("Bird exists with count: %d\n", count)
	} else {
		fmt.Println("Bird does not exist in the map.")
	}
}

// MapWithNestedMaps demonstrates using nested maps
func MapWithNestedMaps() {
	nestedMap := map[string]map[string]int{
		"fruits": {
			"apple":  5,
			"banana": 3,
		},
		"vegetables": {
			"carrot":   4,
			"broccoli": 2,
		},
	}

	// Print the nested map
	fmt.Println("Nested Map:", nestedMap)

	// Access an element in the nested map
	fmt.Println("Apple count in fruits:", nestedMap["fruits"]["apple"])

	// Modify an element in the nested map
	nestedMap["fruits"]["banana"] = 10
	fmt.Println("Modified Nested Map:", nestedMap)

	// Add a new element to a nested map
	nestedMap["fruits"]["grape"] = 6
	fmt.Println("Nested Map after adding grape:", nestedMap)
}
