package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40}

	//b := make([]int, 2, 100)
	//copy(b, a)
	//b[0] = 100
	//b = append(b, 100)

	b := make([]int, len(a[1:3]))
	copy(b, a[1:3])

	fmt.Println(a)
	fmt.Println(b)
}
