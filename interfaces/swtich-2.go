package main

import "fmt"

type User struct {
	//name string
}

func main() {
	checkType(User{})

}

func checkType(i any) {
	switch i.(type) {
	case int:
		fmt.Println("it is an int value")
	case User:
		fmt.Println("this is user")
	default:
		fmt.Println("this is default")

	}
}
