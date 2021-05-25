package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

type checkId struct {
	ID int `validate:"min=1,max=99"`
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/get", echoParamHandler)
	myRouter.HandleFunc("/err", echoBadRequestHandler)
	myRouter.HandleFunc("/url_check", echoCheckIdHandler)

	log.Fatal(http.ListenAndServe(":8080", myRouter))

}

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

func echoCheckIdHandler(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()

	id := r.URL.Query().Get("ID")
	idInt, _ := strconv.Atoi(id)
	query := checkId{
		ID: idInt,
	}

	if err := validate.Struct(query); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	} else {
		fmt.Fprintf(w, "ID：%v\n", idInt)
	}
}
