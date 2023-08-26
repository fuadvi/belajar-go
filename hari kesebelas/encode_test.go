package hari_kesebelas

import (
	"encoding/json"
	"fmt"
	"testing"
)

func ToEncode(data interface{}) string {
	byte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))
	return string(byte)
}

func TestEncode(t *testing.T) {
	ToEncode("fuad")
	ToEncode(1)
	ToEncode(true)
	ToEncode([]string{"teuku", "fuad", "maulana"})
}
