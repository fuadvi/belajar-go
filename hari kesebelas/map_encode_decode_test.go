package hari_kesebelas

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeMap(t *testing.T) {
	stringJson := `{"FirstName":"Teuku","MiddleName":"Fuad","LastName":"Maulana","Hobby":["programmer","basket ball"]}`
	var resultDecode map[string]interface{}
	json.Unmarshal([]byte(string(stringJson)), &resultDecode)

	fmt.Println(resultDecode["FirstName"])
	fmt.Println(resultDecode["MiddleName"])
	fmt.Println(resultDecode["LastName"])
	fmt.Println(resultDecode["Hobby"])
}

func TestEncodeMap(t *testing.T) {
	dataMap := map[string]interface{}{
		"FirstName": "Teuku",
		"LastName":  "Maulana",
	}

	bytes, err := json.Marshal(dataMap)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
