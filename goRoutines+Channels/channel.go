package main

import (
	"fmt"
	"time"
)

func printPing(c chan string) {
	var cont int
	for {
		cont++
		fmt.Println(<-c, " ", cont)
		time.Sleep(time.Second * 1)
	}
}

func sendPing(c chan string) {
	for {
		c <- "Ping"
	}
}

func main() {
	ch := make(chan string)

	go sendPing(ch)
	go printPing(ch)

	var input string
	fmt.Scanln(&input)
	fmt.Println("END...")
}
