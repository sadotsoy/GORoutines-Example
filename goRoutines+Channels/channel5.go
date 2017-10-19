package main

import (
	"fmt"
	"time"
)

func genNumbers(out chan<- int) {
	for x := 1; x <= 5; x++ {
		out <- x
	}
	close(out)
}

func squareNumbers(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printSquaredNumbers(in <-chan int) {
	for x := range in {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	numbers := make(chan int)
	numberSquared := make(chan int)

	go genNumbers(numbers)
	go squareNumbers(numbers, numberSquared)

	printSquaredNumbers(numberSquared)
}
