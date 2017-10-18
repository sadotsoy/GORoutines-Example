package main

import (
	"fmt"
	"time"
)

func main() {
	number := make(chan int)
	squaredNumber := make(chan int)

	go func() {
		for x := 0; ; x++ {
			number <- x
		}
	}()

	go func() {
		for {
			x := <-number
			squaredNumber <- x * x
		}
	}()

	for {
		fmt.Println(<-squaredNumber)
		time.Sleep(1 * time.Second)
	}
}
