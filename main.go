package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/docker", func(w ResponseWriter, req *http.Request) {
		fmt.Printf(w, "<h1>Hello</h1>")
	})
	http.ListenAndServe(":8080", nil)

}
