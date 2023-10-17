package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50}

	// b := a[2:5]

	// b = append(b, 999, 400)

	// fmt.Println(a)
	// fmt.Println(b)

	i := a[1:5]
	fmt.Printf("%p\n", a)

	a = append(a, 600)
	fmt.Printf("%p\n", a)
	i[0] = 999

	fmt.Println(i)
	fmt.Println(a)
	fmt.Println("cap of i", cap(i))
	fmt.Println("len of i", len(i))
	fmt.Println("cap of a", cap(a))
	fmt.Println("len of i", len(a))
}
