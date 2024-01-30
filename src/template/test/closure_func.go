package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", timed(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hello, boy")
	}))
	http.ListenAndServe(":3000", nil)
}
func timed(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		end := time.Now()
		fmt.Println("The request took", end.Sub(start))
	}
}
