package main

import (
	"fmt"
	"sort"
)

type user struct {
	Name string
	Age  int
}

type UserSlice []user

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []user{
		{"maulana", 25},
		{"fuad", 23},
		{"teuku", 22},
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
