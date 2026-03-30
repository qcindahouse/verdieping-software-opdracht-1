package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🎮 Welcome to the Number Guessing Game!")
	fmt.Println("======================================")
	fmt.Println("Rules:")
	fmt.Println("The computer will pick a number between 1 and 100.")
	fmt.Println("You must guess the number within limited attempts.")
	fmt.Println("After each guess, you'll be told if the number is higher or lower.")
	fmt.Println("")

	
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1

	
	fmt.Println("Choose difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (7 chances)")
	fmt.Println("3. Hard (5 chances)")
	fmt.Print("Enter choice (1/2/3): ")

	diffInput, _ := reader.ReadString('\n')
	diffInput = strings.TrimSpace(diffInput)

	var attempts int
	switch diffInput {
	case "1":
		attempts = 10
	case "2":
		attempts = 7
	case "3":
		attempts = 5
	default:
		fmt.Println("Invalid choice. Defaulting to Medium (7 chances).")
		attempts = 7
	}

	fmt.Printf("\nYou have %d attempts to guess the number.\n", attempts)

	
	for i := 1; i <= attempts; i++ {
		fmt.Printf("\nAttempt %d/%d - Enter your guess: ", i, attempts)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("⚠️ Please enter a valid number.")
			i-- 
			continue
		}

		if guess == target {
			fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", i)
			return
		} else if guess < target {
			fmt.Println("Too low! Try a higher number.")
		} else {
			fmt.Println("Too high! Try a lower number.")
		}
	}

	
	fmt.Printf("\n Game Over! You've used all your attempts.\n")
	fmt.Printf("The correct number was: %d\n", target)
}