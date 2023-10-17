package main

import "fmt"

func main() {

	defer fmt.Println(1)
	defer fmt.Println(2)
	panic("some kind of msg")
	var i []int

	i[1] = 1000

}
