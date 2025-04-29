package main

import "fmt"

type Employee struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	a := Employee{1, "first name 1", "last name 1"}

	b := Employee{
		ID:        2,
		FirstName: "first name 2",
		LastName:  "last name 2",
	}

	var c Employee
	c.ID = 3
	c.FirstName = "first name 3"
	c.LastName = "last name 3"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
