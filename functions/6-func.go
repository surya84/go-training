package main

import "fmt"

func main() {
	request(get, "oracle")
	request(post, "sybase")
	request(put, "Mysql")
}

func get(endPoint string) {
	fmt.Println("this is a get method", endPoint)
}
func post(endPoint string) {
	fmt.Println("this is post method", endPoint)
}

func put(endPoint string) {
	fmt.Println("this is put method", endPoint)
}

func request(function func(input string), endPoint string) {
	function(endPoint)
}
