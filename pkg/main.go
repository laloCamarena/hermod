package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello")
}

func handleRequests() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequests()
}
