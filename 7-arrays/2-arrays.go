package main

import "fmt"

func main() {
	var i []int
	//i[1] = 100
	i = append(i, 20, 30)
	fmt.Println(i)
	if i == nil {
		fmt.Println("slice is null", i)

	}
	b := []int{10, 20, 30}
	fmt.Println(b)
}
