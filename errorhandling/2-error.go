package main

import (
	"errors"
	"fmt"
	"log"
)

// var str ="asdefwef"

var ErrStirngCheck = errors.New("string is not found")

func FetchData(str string) (string, error) {
	if str == "" {
		return "", ErrStirngCheck
	}

	return str, nil
}

func main() {
	n, err := FetchData("d")

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("String Found", n)
}
