package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("mod in go")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func greeter() {
	fmt.Println("Hello from greeter")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello</h1>"))
}
