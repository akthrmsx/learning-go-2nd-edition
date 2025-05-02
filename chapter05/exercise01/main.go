package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func add(i int, j int) (int, error) {
	return i + j, nil
}

func sub(i int, j int) (int, error) {
	return i - j, nil
}

func mul(i int, j int) (int, error) {
	return i * j, nil
}

func div(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	} else {
		return i / j, nil
	}
}

var operations = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
		{"10", "/", "0"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}

		left, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		operation, ok := operations[expression[1]]
		if !ok {
			fmt.Println("unsupported operator:", expression[1])
			continue
		}

		right, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := operation(left, right)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(strings.Join(expression, " "), "=", result)
	}
}
