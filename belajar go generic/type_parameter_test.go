package belajar_go_generic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Lenght[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	assert.True(t, true)
}

func TestTypeParameter(t *testing.T) {
	var resultString string = Lenght[string]("fuad")
	assert.Equal(t, "fuad", resultString)

	var resultNumber int = Lenght[int](100)
	assert.Equal(t, 100, resultNumber)
}

func MultipleParam[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleParam(t *testing.T) {
	MultipleParam[string, int]("fuad", 100)
	MultipleParam[int, string](100, "fuad")
}
