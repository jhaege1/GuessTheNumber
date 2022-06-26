package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var min int
	var max int
	var input int

	fmt.Println("Lower bound:")
	fmt.Scanln(&min)
	fmt.Println("Upper bound:")
	fmt.Scanln(&max)

	x := generateRandomNumber(min, max)
	y := calculateGuesses(float64(min), float64(max))

	fmt.Printf("You have %v chances to guess the number!", y)

	count := 0.0

	for count <= y {
		count = count + 1

		fmt.Println("\nGuess the number:")
		fmt.Scanln(&input)

		if input == x {
			fmt.Printf("Congratulations! You did it in %v tries!", count)
			break
		} else if x < input {
			fmt.Println("Too high!")
		} else if x > input {
			fmt.Println("Too low!")
		}

		if count >= y {
			fmt.Printf("The number is %v!", x)
			break
		}
	}
}

func generateRandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(max-min+1) + min

	return number
}

func calculateGuesses(min float64, max float64) float64 {
	numberOfGuesses := math.Round(math.Log(max - min + 1))

	return numberOfGuesses
}
