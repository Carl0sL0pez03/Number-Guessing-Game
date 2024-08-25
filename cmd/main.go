package main

import (
	"fmt"
	"number-guessing-game/internal/game"
	"os"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	for {
		game.Play()
		fmt.Print("Do you want to play again? (y/n): ")
		var response string
		fmt.Scanln(&response)
		if response != "y" {
			fmt.Println("Thank you for playing! Goodbye!")
			os.Exit(0)
		}
	}
}
