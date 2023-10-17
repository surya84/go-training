package main

import (
	"fmt"
	"net/http"
	"project5/handlers"
)

func add(sub func()) {
	fmt.Println("sum")
}

func sub() {
	fmt.Println("sub")
}

func operate(add func()) {

	fmt.Println("sub")
}

func main() {
	http.HandleFunc("/users", handlers.GetUser)
	http.ListenAndServe(":8085", nil)
	sub()
	add(sub)
	operate(add())
}
