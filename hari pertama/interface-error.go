package main

import (
	"errors"
	"fmt"
)

func Pembagi(penjumlah, pembagi int) (int, error) {

	if pembagi == 0 {
		return 0, errors.New("pembagi adalah 0")
	} else {
		return penjumlah / pembagi, nil
	}
}
func main() {
	result, error := Pembagi(100, 0)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}
}
