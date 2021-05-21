package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// /getにアクセスした場合、handleParamsを実行する。
	myRouter.HandleFunc("/get", echoParamHandler)

	log.Fatal(http.ListenAndServe(":8080", myRouter))

}

// http.ResponseWriterは出力先、クエリを受け取る。
func echoParamHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "query：%s\n", r.URL.RawQuery)
}
