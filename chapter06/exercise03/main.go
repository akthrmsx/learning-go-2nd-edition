package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {
	people := make([]Person, 0, 10_000_000)

	for range 10_000_000 {
		people = append(people, MakePerson("a", "b", 1))
	}

	fmt.Println(people)
}
