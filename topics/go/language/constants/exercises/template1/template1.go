// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an untyped and typed constant and display their values.
//
// Multiply two literal constants into a typed variable and display the value.
package main

// Add imports.
import "fmt"

const (
	// Declare a constant named server of kind string and assign a value.
	server = "Phusion"
	// Declare a constant named port of type integer and assign a value.
	port = 32
)

func main() {

	// Display the value of both server and port.
	fmt.Println("server: ", server)
	fmt.Println("port: ", port)
	// Divide a constant of kind integer and kind floating point and
	// assign the result to a variable.
	v := 45 / 2.0

	// Display the value of the variable.
	fmt.Println(v)
}
