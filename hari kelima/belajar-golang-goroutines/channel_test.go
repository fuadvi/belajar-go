package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)

		channel <- "teuku fuad maulana"

		fmt.Println("selesai menunggu chanel mengirim data")
	}()

	fmt.Println(<-channel)
	time.Sleep(5 * time.Second)

	defer close(channel)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "fuad"
}

func TestChannelWithParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	fmt.Println(<-channel)
	fmt.Println("menunggu response")

	time.Sleep(5 * time.Second)

	defer close(channel)
}

func OnlyIn(channel chan<- string) {
	channel <- "maulana"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInAndOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)
	time.Sleep(5 * time.Second)
	fmt.Println("menunggu selesai menggirim channel")
	defer close(channel)
}

func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 3)

	go func() {
		channel <- "teuku"
		channel <- "fuad"
		channel <- "maulana"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 1; i <= 10; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i)
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("ini dari channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("ini dari channel 2 ", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}

		if counter == 2 {
			break
		}
	}

}
