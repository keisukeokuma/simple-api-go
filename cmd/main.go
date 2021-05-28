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
	query := r.URL.RawQuery
	t := time.Now()

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Fprintf(w, "Time：%s\n", err)
		return
	}
	jst_time := t.In(jst)

	if query == "date" {
		const dateFormat = "2006/01/02"
		fmt.Fprintf(w, "Time：%s\n", jst_time.Format(dateFormat))
		return
	} else if query == "time" {
		const timeFormat = "15:04:05"
		fmt.Fprintf(w, "Time：%s\n", jst_time.Format(timeFormat))
		return
	}

	fmt.Fprintf(w, "Time：%s\n", jst_time)
}
