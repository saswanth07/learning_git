package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main2() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting server on http://localhost:8090")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
