package main

import (
	"log"
	"sync"
	"time"
)

var counter = 0

func main() {

	wg := new(sync.WaitGroup)

	var m = sync.Mutex{}

	for i := 0; i < 3; i++ {

		wg.Add(1)

		go func() {

			defer wg.Done()
			m.Lock()
			j := counter

			time.Sleep(time.Millisecond)

			j = j + 1

			counter = j
			m.Unlock()

		}()

	}

	wg.Wait()

	log.Println("counter:", counter)

}
