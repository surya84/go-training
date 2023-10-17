package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg := sync.WaitGroup{}
	//wgWorker :=sync.WaitGroup{}
	wg.Add(2)

	go func() {
		//time.Sleep(2 * time.Second)
		ch1 <- 10
	}()

	go func() {
		ch2 <- 20
	}()

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-ch1:
			fmt.Println("Received 1 : ", m1)

		case m2 := <-ch2:
			fmt.Println("Received 2 :", m2)
		}
	}
}
