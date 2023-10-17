package main

import (
	"fmt"
	"time"
)

func main() {
	go hello() // spawning a goroutine
	fmt.Println("end of the main")

	time.Sleep(2 * time.Second)

}

func hello() {
	time.Sleep(1 * time.Second)
	fmt.Println("hello from the hello func")
}
