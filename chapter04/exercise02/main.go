package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := make([]int, 0, 100)

	for range 100 {
		numbers = append(numbers, rand.Intn(100))
	}

	for _, number := range numbers {
		switch {
		case number%2 == 0 && number%3 == 0:
			fmt.Println("six")
		case number%2 == 0:
			fmt.Println("two")
		case number%3 == 0:
			fmt.Println("three")
		default:
			fmt.Println("never mind")
		}
	}
}
