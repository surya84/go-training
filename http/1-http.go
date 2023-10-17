package main

import "net/http"

func main() {

	http.HandleFunc("/home", handler)
	//start your server
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hi hello"))

}
