package main

import (
	"fmt"
	"log"
	"net/http"

	netflixRouter "github.com/Ashmit-05/mongodbapi/routes"
)

func main() {
	fmt.Println("Server is starting...")
	r := netflixRouter.Router()
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("Listening on port 4000")
}
