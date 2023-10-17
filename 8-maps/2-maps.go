package main

import "fmt"

var dict = make(map[int][]string)

func main() {

	dict[1] = []string{"listening", "playing"}
	dict[2] = []string{"watching", "playing"}
	dict[3] = []string{"seeing", "playing"}

	for k, v := range dict {
		fmt.Println(k, ":", v)
	}

	delete(dict, 3)
	fmt.Println(dict)
	var i [5]int
	var j [5]int
	fmt.Println(i == j)

	m := map[[1]int]interface{}{}

	m[[1]int{1}] = [1]int{1}

	fmt.Println("m:", m)

	id := 5

	search(id)
}

func search(id int) {
	h, k := dict[id]
	if !k {
		fmt.Println("not there", k)
		return
	}

	fmt.Print(h)
}
