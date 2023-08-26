package hari_kesebelas

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Costumer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Hobby      []string
}

func TestEncodeStruct(t *testing.T) {
	coustemer := Costumer{
		FirstName:  "Teuku",
		MiddleName: "Fuad",
		LastName:   "Maulana",
		Hobby:      []string{"programmer", "basket ball"},
	}

	bytes, err := json.Marshal(coustemer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
