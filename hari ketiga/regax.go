package main

import (
	"fmt"
	"regexp"
)

func main() {
	regax := regexp.MustCompile(`f([a-z])d`)
	fmt.Println(regax.MatchString("fud"))
	fmt.Println(regax.MatchString("fUd"))
	fmt.Println(regax.MatchString("fad"))

	// mengambil 2 data yang benar
	fmt.Println(regax.FindAllString("fud fad fat ffd ddd fab", 2))

	// mengambil semua data yang benar
	fmt.Println(regax.FindAllString("fud fad fat ffd ddd fab", -1))
}
