package main

import (
	"fmt"
	"sync"
)

// https://go.dev/ref/spec#Send_statements
// A send on an unbuffered channel can proceed if a receiver is ready. send will block until there is no recv

func main() {

	wg := &sync.WaitGroup{}

	ch := make(chan int, 1) //unbuffered channel
	//ch := make(chan int,3) //buffered channel

	wg.Add(1)

	go func() {
		defer wg.Done()
		ch <- 20 // send will block until no receiver is ready
		ch <- 10
	}()

	x := <-ch //it is a blocking call until we don't recv the value
	xy := <-ch
	fmt.Println(x)
	fmt.Println(xy)
	wg.Wait()
	fmt.Println("end of main")

}
