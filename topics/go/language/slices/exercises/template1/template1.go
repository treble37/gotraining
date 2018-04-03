// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports.
import "fmt"

func main() {

	// Declare a nil slice of integers.
	slice := make([]int, 3)

	// Appends numbers to the slice.
	slice = append(slice, 4)

	// Display each value in the slice.
	for i, value := range slice {
		fmt.Println("i:", i, "value:", value)
	}

	// Declare a slice of strings and populate the slice with names.
	stringSlice := make([]string, 3)
	stringSlice[0] = "joe"
	stringSlice[1] = "jane"
	stringSlice[2] = "jim"

	// Display each index position and slice value.
	for i, value := range stringSlice {
		fmt.Println("i:", i, "value:", value)
	}

	// Take a slice of index 1 and 2 of the slice of strings.
	slicedSlice := stringSlice[1:2]

	// Display each index position and slice values for the new slice.
	fmt.Println("Sliced slices")
	for i, value := range slicedSlice {
		fmt.Println("i:", i, "value:", value)
	}
}
