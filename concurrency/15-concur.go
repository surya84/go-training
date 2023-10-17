package main

import (
	"fmt"

	"io"

	"log"

	"net/http"

	"sync"
)

var url1 = []string{

	`https://pkg.go.dev/`,

	`https://github.com/`,

	`abc.com/1234`,
}

type response1 struct {
	url  string
	resp *http.Response
	err  error
}

// var wg = &sync.WaitGroup{}

var wgWorker = sync.WaitGroup{}

func main() {

	//create a struct which stores the response, error and the url name

	//
	respChan1 := make(chan response1)
	go doGetRequest1(url1, respChan1)
	fetchResults1(respChan1)

}

func doGetRequest1(urls []string, respChan1 chan response1) {

	for _, url := range urls {

		wgWorker.Add(1)
		go func(url string) {

			defer wgWorker.Done()
			resp, err := http.Get(url)

			r1 := response1{
				url:  url,
				resp: resp,
				err:  err,
			}

			respChan1 <- r1 //sending the resp struct to respCh

		}(url)

	}

	wgWorker.Wait()

	close(respChan1)

	//fetchResults1(respChan)

}

func fetchResults1(respChan1 chan response1) {

	// Loop through every element in respChan

	for r := range respChan1 {

		// If there is an error with the response object, log it and continue with the next iteration

		if r.err != nil {

			log.Println(r.err)

			continue

		}

		// Attempt to read the entire response body

		bytes, err := io.ReadAll(r.resp.Body)

		// If an error occurred during reading, log the error and continue with the next iteration

		if err != nil {

			log.Println(err)

			continue

		}

		// Ensuring that the response body is eventually closed, to prevent potential memory leak

		func() {

			defer r.resp.Body.Close()

		}()

		// If the status code indicates an unsuccessful response (i.e., is greater than 299)

		// log the status code and the body content, then move on to the next iteration

		if r.resp.StatusCode > 299 {

			log.Printf("Response failed with status code: %d and\nbody: %s\n", r.resp.StatusCode, bytes)

			continue

		}

		// If the response is successful, print out the URL associated with it and its status

		fmt.Println(r.url, r.resp.Status)

	}

	//some more work

}
