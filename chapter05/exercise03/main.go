package main

import "fmt"

func prefixer(prefix string) func(string) string {
	return func(body string) string {
		return fmt.Sprintf("%s %s", prefix, body)
	}
}

func main() {
	hello := prefixer("hello")

	fmt.Println(hello("a"))
	fmt.Println(hello("b"))
}
