# Number Guessing Game

This is a simple command-line based number guessing game built in Go. The computer randomly selects a number between 1 and 100, and the user has to guess the number within a limited number of attempts based on the selected difficulty level.

## How to Play

1. Clone the repository.
2. Build the game using the provided `Makefile`.
3. Run the game from the command line.
4. Select a difficulty level and start guessing the number.
5. The game will provide hints whether the number is higher or lower than your guess.
6. The game ends when you guess the number correctly or run out of chances.

## Features

- Selectable difficulty levels (Easy, Medium, Hard)
- Feedback on whether the guess is too high or too low
- Multiple rounds of play
- Track time taken to guess the number

## Installation

```bash
git clone https://github.com/your-repo/number-guessing-game.git
cd number-guessing-game
make build
```

## Usage

Run the CLI with a GitHub username as an argument:

```sh
make run
```

# Challenge roadmap

- https://roadmap.sh/projects/number-guessing-game