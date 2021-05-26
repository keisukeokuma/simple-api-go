package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

type checkId struct {
	ID   int    `validate:"min=1,max=99"`
	Name string `validate:"required"`
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/get", echoParamHandler)
	myRouter.HandleFunc("/err", echoBadRequestHandler)
	myRouter.HandleFunc("/url_check", echoCheckIdHandler)
	myRouter.HandleFunc("/time", echoTimeHandler)

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

	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Print("bad param.")
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	query := checkId{
		ID:   idInt,
		Name: name,
	}

	if err := validate.Struct(query); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "ID：%v\nName:%v", idInt, name)
}

func echoTimeHandler(w http.ResponseWriter, r *http.Request) {
	// query := r.URL.RawQuery
	// time := time.Now()
	layout := "Jan 2, 2006 at 3:04pm (MST)"
	value := "Feb 3, 2013 at 7:54pm (PST)"
	t, _ := time.Parse(layout, value)

	// if query == "" {

	// }
	fmt.Fprintf(w, "Time：%s\n", t)
}
