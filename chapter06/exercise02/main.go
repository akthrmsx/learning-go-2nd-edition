package main

import "fmt"

func UpdateSlice(xs []string, x string) {
	xs[len(xs)-1] = x
	fmt.Println("in UpdateSlice:", xs)
}

func GrowSlice(xs []string, x string) {
	xs = append(xs, x)
	fmt.Println("in GrowSlice:", xs)
}

func main() {
	xs := []string{"a", "b", "c"}

	UpdateSlice(xs, "d")
	fmt.Println("in main after UpdateSlice:", xs)

	GrowSlice(xs, "e")
	fmt.Println("in main after GrowSlice:", xs)
}
