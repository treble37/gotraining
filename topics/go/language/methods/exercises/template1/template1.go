// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

// Add imports.
import (
	"fmt"
	"math/rand"
)

// Declare a struct that represents a ball player.
// Include field called name, atBats and hits.
type player struct {
	name   string
	atBats int
	hits   int
}

// Declare a method that calculates the batting average for a player.
func (p *player) average() float64 {
	if p.atBats == 0.0 {
		return 0.0
	}
	return float64(p.hits) / float64(p.atBats)
}

func main() {

	// Create a slice of players and populate each player
	// with field values.
	players := make([]player, 2)
	players[0] = player{
		name:   "John0",
		atBats: rand.Intn(100),
		hits:   rand.Intn(100),
	}
	players[1] = player{
		name:   "John1",
		atBats: rand.Intn(100),
		hits:   rand.Intn(100),
	}
	// Display the batting average for each player in the slice.
	for _, player := range players {
		fmt.Printf("hits[%d] atBats[%d] average[%f]\n", player.hits, player.atBats, player.average())
	}
}
