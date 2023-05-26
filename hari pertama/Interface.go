package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) (int, error) {
	return fmt.Println("hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	var users Person
	users.Name = "fuad"

	SayHello(users)
}
