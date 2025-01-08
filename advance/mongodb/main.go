package main

import (
	"fmt"
	"log"
	"net/http"

	"mongodb/router"
)

func main() {

	fmt.Println("MongoDB API")

	r := router.Router()
	fmt.Println("Server Started")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")
}
