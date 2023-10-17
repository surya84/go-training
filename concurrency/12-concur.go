package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg2 := sync.WaitGroup{}
	// wg2.Add(1)
	wg.Add(1)
	go func() {
		// wg2.Done()
		wg.Done()
		for i := 0; i <= 10; i++ {
			wg2.Add(1)
			go func(i int) {
				defer wg2.Done()
				ch <- i
			}(i)
		}
		wg2.Wait()
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range ch {

			fmt.Println("Recieved : ", v)

		}
	}()
	wg.Wait()
}
