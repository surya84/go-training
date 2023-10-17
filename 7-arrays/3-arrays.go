package main

import "fmt"

func main() {
	//var i[]int
	i := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(i[1:4])
	fmt.Println(i[:5])
	fmt.Println(i[1:])

	fmt.Println("slice")
	a := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(a)
	b := a[1:4]
	fmt.Println(b)
	b[0] = 777
	fmt.Println(b)
	fmt.Println(a)
}
