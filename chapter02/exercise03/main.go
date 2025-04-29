package main

import "fmt"

func main() {
	var a byte = 255
	var b int32 = 2147483647
	var c uint64 = 18446744073709551615

	a += 1
	b += 1
	c += 1

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
