// Made by Mohamad
// Made on May 2023
//
// This program counts up to a number of the user's choice

package main

import (
	"fmt"
)

func main() {
	// Declare variables
	var numberGoal int
	var counter int

	// Ask user for input
	fmt.Println("Enter a number to count up to: ")
	fmt.Scanln(&numberGoal)
	fmt.Println("") // Blank line to seperate input from output

	// Count
	for counter = 0; counter <= numberGoal; counter++ {
		fmt.Println(counter)
	}

	fmt.Println("\nDone.")
}
