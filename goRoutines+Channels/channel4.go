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
		for x := range number {
			squaredNumber <- x * x
		}
		close(squaredNumber)
	}()

	for x := range squaredNumber {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}
