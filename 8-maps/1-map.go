package main

import (
	"fmt"
)

func main() {
	dictionary := make(map[string]string) //make(map[string]string,10)

	dictionary["up"] = "above"
	dictionary["below"] = "down"

	fmt.Println(dictionary["up"])

	for key := range dictionary {
		fmt.Println(key)
	}

	//value
	for _, value := range dictionary {
		fmt.Println(value)
	}

	fmt.Println(len(dictionary))
	fmt.Println(dictionary)
}
