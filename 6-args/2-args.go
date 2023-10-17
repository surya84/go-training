package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	greet()
	fmt.Println("doiing some more work in main func")
	fmt.Println("end of main")

}

func greet() {
	data := os.Args[1:]
	if len(data) != 3 {
		log.Println("please provide name , age , marks")
		os.Exit(1)
		//return // stops the exec of the current func
	}

	name := data[0]
	ageString := data[1]
	age, err := strconv.Atoi(ageString)
	if err != nil {
		log.Println(err)
		//return
		os.Exit(1)
	}

	marksString := data[2]
	marks, err := strconv.Atoi(marksString)

	if err != nil {
		// there is some kind of error
		log.Println(err)
		//return
		os.Exit(1)
	}
	fmt.Println(name, age, marks)
}
