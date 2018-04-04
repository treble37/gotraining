// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

// Add imports.
import "fmt"

func main() {

	// Declare and make a map of integer type values.
	hollywood := make(map[string]int)

	// Initialize some data into the map.
	hollywood["Beverly Hills"] = 90201
	hollywood["Santa Monica"] = 90291
	hollywood["Westwood"] = 96137
	// Display each key/value pair.
	for k, v := range hollywood {
		fmt.Printf("key[%s] value[%d]\n", k, v)
	}
}
