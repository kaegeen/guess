package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Initialize random seed based on current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	target := rand.Intn(100) + 1
	var guess int
	var attempts int

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("Can you guess what it is?")

	// Game loop
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)
		attempts++

		// Check the user's guess
		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You guessed the number %d in %d attempts.\n", target, attempts)
			break
		}
	}
}
