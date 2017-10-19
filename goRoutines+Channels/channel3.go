package main

import (
	"fmt"
	"time"
)

func main() {
	number := make(chan int)
	squaredNumber := make(chan int)

	go func() {
		for x := 1; x <= 5; x++ {
			number <- x
		}
		close(number)
	}()

	go func() {
		for {
			x, ok := <-number
			if !ok {
				break
			}
			squaredNumber <- x * x
		}
		close(squaredNumber)
	}()

	for {
		x, ok := <-squaredNumber
		if !ok {
			break
		}
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}
