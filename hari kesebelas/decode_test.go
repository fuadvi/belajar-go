package hari_kesebelas

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	costumer := Costumer{
		FirstName:  "Teuku",
		MiddleName: "Fuad",
		LastName:   "Maulana",
		Hobby:      []string{"programmer", "basket ball"},
	}
	stringJson := ToEncode(costumer)

	resultCostumer := &Costumer{}

	json.Unmarshal([]byte(stringJson), resultCostumer)

	fmt.Println(resultCostumer.FirstName)
	fmt.Println(resultCostumer.MiddleName)
	fmt.Println(resultCostumer.LastName)

}
