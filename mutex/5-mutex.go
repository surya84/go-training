package main

import (
	"fmt"
	"sync"
)

var wg3 = sync.WaitGroup{}

type Car struct {
	driver int
	rw     sync.RWMutex
}

func (c *Car) getCarDriver() {
	c.rw.RLock()
	defer c.rw.RUnlock()
	fmt.Println("Driver :", c.driver)
}

func main() {
	c := Car{
		driver :5,
	}
	
}
