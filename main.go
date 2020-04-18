package main

import (
	"cd_rental/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/collection/", controllers.GetCollection)
	http.HandleFunc("/collection", controllers.PostCollection)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
