package main

import "fmt"

type userSlice struct {
	name  string
	email string
}

func main() {
	u := []userSlice{
		{
			name: "surya",
		},
		{
			email: "surya@gmail.com",
		},
	}

	fmt.Println(u)
}
