// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"hogehoge"
// 	"github.com/go-playground/validator"
// 	"github.com/gorilla/mux"
// )

// type ErrUrl string

// func (url ErrUrl) Error() string {
// 	return fmt.Sprintf("%v", string(url))
// }

// type Foo struct {
// 	ID   `validate:"-"`
// 	Name `validate:"required,max=255"`
// 	Kana `validate:"max=255"`
// }

// func init() {
// 	validate := validator.New()
// }

// func validate(foo *Foo) error {
// 	return validate.Struct(foo)
// }

// // http.ResponseWriterは出力先、クエリを受け取る
// func echoHandle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "query：%s\n", r.URL.RawQuery)
// }

// func errHandle(w http.ResponseWriter, r *http.Request) {
// 	url := r.URL.RawQuery
// 	errCheck(url)
// }

// func errCheck(url string) error {
// 	if url == "" {
// 		return ErrUrl(url)
// 	} else {
// 		return nil
// 	}
// }

// func main() {
// 	myRouter := mux.NewRouter().StrictSlash(true)

// 	// /getにアクセスした場合、handleParamsを実行する
// 	myRouter.HandleFunc("/get", echoHandle)
// 	myRouter.HandleFunc("/err", errHandle)

// 	log.Fatal(http.ListenAndServe(":8080", myRouter))

// }
