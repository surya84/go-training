package main

import (
	"fmt"
	"strconv"
)

func Name(a, b string) string {
	return a + b
}
func square(a, b string) (int, error) {

	//var sum int

	i, err := strconv.Atoi(a)
	if err != nil {
		return i, err
	}

	j, err := strconv.Atoi(b)
	if err != nil {
		return j, err
	}

	return i + j, nil

}

func main() {
	//fmt.Println(Name("surya", "teja"))
	b, err := square("2fe", "3")

	if err == nil {
		fmt.Println("no error", b)
	} else {
		fmt.Println("error", err)
	}
}
