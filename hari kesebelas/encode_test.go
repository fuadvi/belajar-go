package hari_kesebelas

import (
	"encoding/json"
	"fmt"
	"testing"
)

func toEncode(data interface{}) {
	byte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))
}

func TestEncode(t *testing.T) {
	toEncode("fuad")
	toEncode(1)
	toEncode(true)
	toEncode([]string{"teuku", "fuad", "maulana"})
}
