package main

import (
	"log"
	"net/http"
)

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
