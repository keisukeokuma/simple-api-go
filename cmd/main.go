package main

import (
	"fmt"
	"log"
	"net/http"

	"simple-api-go/internal/hoge"

	"github.com/gorilla/mux"
)

// type User struct {
// 	FamilyName string `validate:"required"` // 必須項目
// }

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/get", echoParamHandler)
	myRouter.HandleFunc("/err", echoBadRequestHandler)
	// myRouter.HandleFunc("/validate", echoValidateHandler)

	hoge.Hoge()
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

// func echoValidateHandler(w http.ResponseWriter, r *http.Request) {
// 	validate = validator.New()
// 	// err := validate.Struct(r.URL.RawQuery)
// 	query := r.URL.RawQuery
// 	err := validate(query)
// 	if query != "" {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 	} else {
// 		fmt.Fprintf(w, "query：%s\n", query)
// 	}
// }

// func init() {
// 	validate := validator.New()
// }

// func validate(query *User) error {
// 	return validate.Struct(query)
// }
