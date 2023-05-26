package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `require:"true" max:"10"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("require") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}

	return true
}

func main() {
	sampel := Sample{"fuad"}

	var data reflect.Type = reflect.TypeOf(sampel)

	fmt.Println(data.Field(0).Name)
	fmt.Println(data.NumField())
	fmt.Println(data.Field(0))
	sampel.Name = ""
	fmt.Println(data.Field(0).Tag.Get("require"))
	fmt.Println(reflect.ValueOf(sampel).Field(0).Interface())
	fmt.Println(IsValid(sampel))
}
