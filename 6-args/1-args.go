package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("before slicing")
	//fmt.Printf("%T", os.Args)
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])

	fmt.Println("this is from after slicing")
	a := os.Args[1:]
	fmt.Println(a)

}
