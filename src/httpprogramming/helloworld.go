package main

import (
	"fmt"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello world, %s", req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}
