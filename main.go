package main

import (
	"fmt"
	"time"
)

func main() {

	var input int

	locations := []string{"Gotham city", "Metroplis", "Queens", "Little Land", "Vice City"}
	superPowers := []string{"Laser eyes", "Flying", "Invisibility", "Power Absobtion", "Fire", "Freezing", "SuperHuman Speed"}

	fmt.Println("Welcome to the world superheroes with superpowers")
	fmt.Println("Your choice will unveil your superpower:")

	//Using current time as a seed
	seed := time.Now().UnixNano()

	//User input

	fmt.Println("Choose a number.")
	fmt.Scan(&input)

	//Modifying the seed with user input
	seed = (seed + int64(input)) % 100

	//Selecting the strings based on modified seed

	choice1 := locations[seed%int64(len(locations))]
	choice2 := superPowers[seed%int64(len(superPowers))]

	fmt.Printf("Your superpower is %s in %s.\n", choice2, choice1)
}
