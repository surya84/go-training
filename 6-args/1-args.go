package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

	name := data[0]
	ageString := data[1]
	age, err := strconv.Atoi(ageString)

	var err error

	if err != nil {
		log.Println(err)
	}

	fmt.Println(err)

}
