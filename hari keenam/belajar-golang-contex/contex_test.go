package main

import (
	"context"
	"fmt"
	"testing"
)

func TestContex(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestWIthValue(t *testing.T) {
	background := context.Background()

	backgroundA := context.WithValue(background, "a", "A")
	backgroundB := context.WithValue(background, "b", "B")
	backgroundC := context.WithValue(backgroundA, "c", "C")
	backgroundD := context.WithValue(backgroundA, "d", "D")
	backgroundE := context.WithValue(backgroundB, "e", "E")
	backgroundF := context.WithValue(backgroundC, "f", "F")

	fmt.Println(background)
	fmt.Println(backgroundA)
	fmt.Println(backgroundB)
	fmt.Println(backgroundC)
	fmt.Println(backgroundD)
	fmt.Println(backgroundE)
	fmt.Println(backgroundF)
}
