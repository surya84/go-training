package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type reqKey int

const RequestIDKey reqKey = 123

func main() {
	http.HandleFunc("/home", RequestIdMid(LoggingMid(Middle3(homePage))))
	http.ListenAndServe(":8084", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {

	log.Println("In home Page handler")
	fmt.Fprintln(w, "this is my home page")

}

func RequestIdMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid.NewString()
		ctx := r.Context()
		ctx = context.WithValue(ctx, RequestIDKey, uuid)
		next(w, r.WithContext(ctx))

	}
}
func LoggingMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqId, ok := ctx.Value(RequestIDKey).(string)
		if !ok {
			reqId = "Unknown"
		}
		log.Printf("%s : started   : %s %s ",
			reqId,
			r.Method, r.URL.Path)
		defer log.Println("completed")
		next(w, r)
	}

}

func Middle3(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqId, ok := ctx.Value(RequestIDKey).(string)
		log.Println(reqId, ok)
	}
}
