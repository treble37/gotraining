// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var i int
	var j float64
	var pi float32
	pi = 3.14

	// Display the value of those variables.
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(pi)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	v := "Hallmark Labs"
	w := 32

	// Display the value of those variables.
	fmt.Println(v)
	fmt.Println(w)

	// Perform a type conversion.
	x := float64(w)

	// Display the value of that variable.
	//fmt.Printf("%T", x)
	fmt.Println(x)
}
