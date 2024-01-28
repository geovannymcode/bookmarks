package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello World!")
	if err != nil {
		log.Println("Error processing the request")
	}
}
