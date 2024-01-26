package main

import (
	"fmt"
	"time"
)

func main() {
	// variables
	var input int

	superheroes := []string{"Batman", "Superman", "Spiderman", "Ironman", "Wonder Women"}
	villans := []string{"The Joker", "Loki", "Thanos", "Darth Vader", "Voldemort"}
	quests := []string{"catch", "destroy", "kill", "save", "kidnap"}

	// introduction
	fmt.Println("Hey...Hey you, do you want to become a superhero?")
	fmt.Println("I will give you a quest task and a villain to defeat...")

	// User input
	fmt.Print("\nEnter any 1 digit number: ")
	fmt.Scan(&input)

	// Selecting the strings based on modified user input
	// generating seed
	seed := randomNum(input)
	// hero selection
	choice1 := selector(seed, superheroes)
	// renewing seed
	seed = randomNum(input)
	// villan selection
	choice2 := selector(seed, villans)
	// renewing seed
	seed = randomNum(input)
	// task selection
	choice3 := selector(seed, quests)

	// quest assigment
	fmt.Printf("\nYou are %s and your quest is to %s %s\n", choice1, choice3, choice2)
}

func randomNum(input int) int64 {
	/*function to produce a "random" number based on a single digit integer.
	  It takes an integer as argument and return another integers based on
	  the start of the unix time in nanoseconds
	*/

	// getting unix time
	seed := time.Now().UnixMicro()
	// incorporating user input
	seed += int64(input)
	// reducind the number to a antegers less than 100
	seed %= 100

	return seed

}

func selector(seed int64, array []string) string {
	/* Function to select an element from an array
	   based on randomized user input. The seed being a < 100
	   is reduced to a number no larger than the size of the array.
	*/

	// the seed being a 2 digit number < 100 is reduced to a single digit
	selection := array[seed%int64(len(array))]

	return selection
}
