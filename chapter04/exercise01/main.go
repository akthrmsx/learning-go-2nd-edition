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

	fmt.Println(numbers)
}
