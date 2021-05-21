package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// /getにアクセスした場合、handleParamsを実行する
	myRouter.HandleFunc("/get", echoParamHandler)
	myRouter.HandleFunc("/err", echoBadRequestHandler)

	log.Fatal(http.ListenAndServe(":8080", myRouter))

}

// http.ResponseWriterは出力先、クエリを受け取る
func echoParamHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RawQuery
	fmt.Fprintf(w, "query：%s\n", query)
}

func echoBadRequestHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RawQuery
	if query == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
	} else {
		fmt.Fprintf(w, "query：%s\n", query)
	}

}
