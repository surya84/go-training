package main

import (
	"fmt"
	"sync"
)

var myMap = make(map[int]int)
var wg sync.WaitGroup

func main() {
	var sm sync.Map

// Store key-value pairs in the sync.Map
sm.Store("color", "green")
	var m = sync.Mutex{}

	//defer m.Unlock()
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(n int) {
			//m.Lock()
			defer wg.Done()
			m.Lock()
			defer m.Unlock()
			myMap[n] = n * n

		}(i)
	}
	wg.Wait()
	fmt.Println("Map : ", myMap)
	fmt.Println("End of program")
}
