package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/",homePage).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000",r))
}

func greeter()  {
	fmt.Println("Hi there")
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte(`<h1>Yo Ho a pirate's life for me!`))
}