package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	UserId  string
	Hobbies []string
}

func main() {

	http.HandleFunc("/home", handler1)
	//start your server
	http.ListenAndServe(":8080", nil)
}

func handler1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u1 := User{
		UserId:  "surya",
		Hobbies: []string{"playing"},
	}

	j, err := json.Marshal(u1)

	if err != nil {
		w.Write([]byte("some erorr"))
		return
	}
	w.WriteHeader(http.StatusBadGateway)
	w.Write(j)

}
