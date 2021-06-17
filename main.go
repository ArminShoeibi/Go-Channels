package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for {
			channel1 <- "Every 500 Milliseconds"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			channel2 <- "Every 2 Seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case chan1 := <-channel1:
			fmt.Println(chan1)
		case chan2 := <-channel2:
			fmt.Println(chan2)
		}
	}

}
