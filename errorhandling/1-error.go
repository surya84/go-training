package main

import (
	"errors"
	"fmt"
	"log"
)

var user = make(map[int]string)

// Error string should not be capitalized or end with punctuation mark
// Error strings should not be capitalized (unless beginning with an exported name, a proper noun or an acronym) and should not end with punctuation. This is because error strings usually appear within other context before being printed to the user.

var ErrRecordNotFound = errors.New("record not found")

func FetchRecord(id int) (string, error) {

	name, ok := user[id]
	if !ok {
		return "", ErrRecordNotFound
	}
	return name, nil
}

func main() {

	n, err := FetchRecord(10)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(n)
}
