package main

import "fmt"

func main() {
	a := make([]int, 2, 3)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)
	fmt.Printf("%p", a)
	a = append(a, 100, 200)
	fmt.Println(a)
	fmt.Println()
	fmt.Printf("%p", a)
}
