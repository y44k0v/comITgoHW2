package main

import (
	"fmt"
	"time"
)

func main() {
	var input int
	superheroes := []string{"Batman", "Superman", "Spiderman", "Ironman", "Wonder Women"}
	villans := []string{"The Joker", "Loki", "Thanos", "Darth Vader", "Voldemort"}
	Quests := []string{"Catch", "....", "kill", "save", "live", "go", "rage"}

	fmt.Println("Hey...Hey you, do you want to become a superhero?")
	fmt.Println("I will give you a quest and a villain to defeat:")

	// Using current time as a seed
	seed := time.Now().UnixNano()

	// User input
	fmt.Println("Choose a number:")
	fmt.Scan(&input)

	// Modifying the seed with user input
	seed = (seed + int64(input)) % 100

	// Selecting the strings based on modified seed
	choice1 := superheroes[seed%int64(len(superheroes))]
	choice2 := villans[seed%int64(len(villans))]
	choice3 := Quests[seed%int64(len(Quests))]

	fmt.Printf("You are  %s and your quest is to  %s  %s\n", choice1, choice3, choice2)
}
