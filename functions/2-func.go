package main

import "fmt"

func Operate(op func(x, y int) int) {
	fmt.Println(op(12, 3))
}

func main() {
	func add(a, b int) int {
		return a + b
	}

	fmt.Println(add(10, 20)) //30

	Operate()
}
