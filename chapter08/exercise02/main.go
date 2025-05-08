package main

import (
	"fmt"
	"strconv"
)

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintInt int

func (i PrintInt) String() string {
	return strconv.Itoa(int(i))
}

type PrintFloat float64

func (f PrintFloat) String() string {
	return fmt.Sprintf("%f", f)
}

func Println[T Printable](t T) {
	fmt.Println(t)
}

func main() {
	var i PrintInt = 10
	Println(i)

	var f PrintFloat = 10.123
	Println(f)
}
