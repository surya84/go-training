package main

import "fmt"

func main() {
	operate(sub)
}

func sub(x, y int) int {
	return x - y
}

func

func operate(op func(int, int) int) {
	fmt.Println(op(20, 10))
}
