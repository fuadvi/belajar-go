package belajar_golang_goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("total cpu ", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine running", totalGoroutine)
}
