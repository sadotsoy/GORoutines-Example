package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg is used for indicate to the program to wait the ending of the GoRoutines
var wg sync.WaitGroup

func main() {
	//We add 2 to wg to wait the ending our 2 GoRoutines
	wg.Add(2)
	fmt.Println("Starting the GoRoutines")
	//We launch the GoRoutine with the tag "A"
	go printString("A")
	//We launch the GoRoutine with the tag "B"
	go printString("B")
	// Wait for the GoRoutines end
	fmt.Println("Waiting the end...")
	wg.Wait()
	//When the GoRoutines are over then print THE END!
	fmt.Println("THE END!>>>")
}

func printString(tag string) {
	//We call the function Done() from wg to say that is already over
	defer wg.Done()
	//Random waiting.
	for quantity := 1; quantity <= 10; quantity++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Quantity: %d of %s\n", quantity, tag)
	}
}
