package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	man := Man{"fuad"}
	man.married()

	fmt.Println(man.Name)
}
