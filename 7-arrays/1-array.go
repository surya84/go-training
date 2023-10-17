package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b[4])
	c := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(len(c))
}
