package main

import (
	"log"
	"net/http"
)


func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}