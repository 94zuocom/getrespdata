package main

import (
	"fmt"
	"log"
	"net/http"
	"getrespdata/p" 
)

func main() {
	http.HandleFunc("/r", p.RequestHandler)
	http.HandleFunc("/time", p.TimeHandler)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

