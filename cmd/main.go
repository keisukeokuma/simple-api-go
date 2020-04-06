package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/echo", echoHandler)

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

// echoHandler ...
func echoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK!")
}
