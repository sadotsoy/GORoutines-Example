package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Println("Insert the number of interaction for Fibonacci")
	fmt.Scanf("%d", &n)
	go animation(100 * time.Millisecond)
	result := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, result)
}

func animation(delay time.Duration) {
	for {
		for _, r := range `\|/-` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
