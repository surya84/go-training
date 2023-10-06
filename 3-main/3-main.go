package main

import (
	"log"
	"os"
)

func main() {

	data := os.Args[1:]
	if len(data) != 3 {
		log.Println("Provide values")
		return
	}
}
