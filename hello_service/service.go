package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is my website!")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

	})
}
