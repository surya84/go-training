package main

import "fmt"

func main() {
	operate(add)

}
func operate(op func(x, y int) int) {
	fmt.Println(op(10, 20))
}

func add(a, b int) int {
	return a + b
}
