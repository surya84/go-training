package main

import "fmt"

type money int

type user struct {
	name  string
	age   int
	marks []int
}

func main() {
	var rupee money = 10
	fmt.Printf("%T ", rupee)
	i := 120
	rupee = money(i)
	fmt.Println(rupee)

	u := user{
		name:  "surya",
		age:   22,
		marks: []int{1, 2, 3},
	}

	fmt.Printf("%+v", u)
	fmt.Println()
	fmt.Println(u)
	fmt.Println()
	fmt.Print(u.name)
	fmt.Print(u.age)
	fmt.Println()
	fmt.Printf("%#v", u)

}
