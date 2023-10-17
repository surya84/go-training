package main

import (
	"net/http"
	"project4/handlers"
)

type User struct {
	UserId  string
	Hobbies []string
}

func main() {

	http.HandleFunc("/home", handlers.Home)
	//start your server

	panic(http.ListenAndServe(":8082", nil))

}
