package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("teuku")
	data.PushBack("fuad")
	data.PushBack("maulana")
	data.PushBack(1000)

	// menampilkan isi double list

	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println("========================")
	//revers
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
