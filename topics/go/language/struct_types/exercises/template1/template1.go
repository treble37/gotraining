// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.

import "fmt"

// Add user type and provide comment.

type user struct {
	email string
	name  string
}

func main() {

	// Declare variable of type user and init using a struct literal.
	u := user{email: "john@hallmarklabs.com", name: "John Smith"}

	// Display the field values.
	fmt.Println(u.email)
	fmt.Println(u.name)

	// Declare a variable using an anonymous struct.
	v := struct {
		brand string
		model string
	}{
		brand: "Ford",
		model: "Ranger",
	}
	// Display the field values.
	fmt.Println(v.brand)
}
