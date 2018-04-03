// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

// Add imports.
import "fmt"

func main() {

	type user struct {
		email string
		name  string
	}
	var users = []user{
		user{
			email: "john@hallmarklabs.com",
			name:  "john",
		},
		user{
			email: "jane@hallmarklabs.com",
			name:  "jane",
		},
	}
	// Declare an array of 5 strings set to its zero value.
	var fiveZero [5]string

	// Declare an array of 5 strings and pre-populate it with names.
	// Assign the populated array to the array of zero values.
	fivePopulated := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	// Iterate over the first array declared.
	// Display the string value and address of each element.
	for i, value := range fiveZero {
		fmt.Println("i ", i, "value: ", value)
	}
	for i, value := range fivePopulated {
		fmt.Println("i ", i, "value: ", value)
	}
	for i, value := range users {
		fmt.Println("user i ", i, "value: ", value.email)
	}
}
