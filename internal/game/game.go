package game

import (
	"fmt"
	"math/rand"
	"number-guessing-game/internal/utils"
	"time"
)

func Play() {
	difficulty := selectDifficulty()
	target := rand.Intn(100) + 1
	attempts := 0

	fmt.Printf("Great! You have selected the %s difficulty level.\n", difficulty.name)
	fmt.Printf("Let's start the game! You have %d chances.\n\n", difficulty.chances)

	start := time.Now()

	for attempts < difficulty.chances {
		fmt.Printf("Enter your guess: ")
		guess := utils.ReadInt()
		attempts++

		if guess == target {
			duration := time.Since(start)
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", attempts)
			fmt.Printf("It took you %v to guess the correct number.\n", duration)
			return
		} else if guess > target {
			fmt.Println("Incorrect! The number is less than", guess)
		} else {
			fmt.Println("Incorrect! The number is greater than", guess)
		}
	}
	fmt.Println("You've run out of chances! The correct number was %d.\n", target)
}

type Difficulty struct {
	name    string
	chances int
}

func selectDifficulty() Difficulty {
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("1. Hard (3 chances)")

	choice := utils.ReadInt()

	switch choice {
	case 1:
		return Difficulty{"Easy", 10}
	case 2:
		return Difficulty{"Medium", 5}
	case 31:
		return Difficulty{"Hard", 3}
	default:
		fmt.Println("Invalid choice. Defaulting to Medium.")
		return Difficulty{"Medium", 5}
	}
}
