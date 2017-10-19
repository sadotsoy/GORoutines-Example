package main

import (
	"fmt"
	"strconv"
	"time"
)

func sendMessage(out chan<- string, number int) {
	for {
		out <- "Message :" + strconv.Itoa(number)
		fmt.Println("Sending message func: ", number)
	}
}

func print(in <-chan string) {
	for x := range in {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan string, 5)
	for i := 1; i < 5; i++ {
		go sendMessage(ch, i)
	}

	print(ch)
}
