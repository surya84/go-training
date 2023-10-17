package main

import (
	"fmt"
	"net/http"
)

type HttpStruct struct {
	Response *http.Response
	Err      error
	Url      string
}

func main() {
	resp, err := http.Get("https://gobyexample.com/select")
	// fmt.Println("response : ", resp)
	// fmt.Println("err", err)

	h := HttpStruct{
		Response: resp,
		Err:      err,
		Url:      "https://gobyexample.com/select",
	}

	fmt.Println("Response : ", h.Response)
	fmt.Println("Error : ", h.Err)
	fmt.Println("Url : ", h.Url)

}
