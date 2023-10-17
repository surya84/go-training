package main

import (
	"fmt"

	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	// counter to add number of goroutines
	// that we are starting or spinning up
	wg.Add(2)
	go Hello(wg) // spawning a goroutine
	go hi(wg)

	fmt.Println("end of the main")

	wg.Wait() //block until the counter is reset to 0

	//bad way of doing it
	//time.Sleep(2 * time.Second)
}

func Hello(wg *sync.WaitGroup) {
	defer wg.Done() //decrement the counter
	fmt.Println("hello from the hello func")
}

func hi(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hi from the hi func")
}
