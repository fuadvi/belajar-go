package main

import "fmt"

func Name() interface{} {
	return "fuad"
}

func main() {
	//var name interface{} = Name()
	//var nameString string = name.(string)
	//fmt.Println(nameString)

	// handling panic

	var name interface{} = Name()
	switch value := name.(type) {
	case string:
		fmt.Println("value", value, "is string")
	case int:
		fmt.Println("value", value, "is integer")
	default:
		fmt.Println("unknown type")
	}
}
