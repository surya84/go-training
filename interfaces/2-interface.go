package main

import (
	"fmt"
	"log"
)

type user struct {
	name string
}

// Write implements io.Writer.
func (user) Write(p []byte) (int, error) {
	fmt.Println("writer method")
	return 0, nil
}

func main() {
	u := user{name: "ajay"}
	l := log.New(u, "sales:", log.Lshortfile)
	l.Println()
}
