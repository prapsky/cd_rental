package main

import (
	"cd_rental/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/collection", controllers.PostCollection)
	http.HandleFunc("/collection/", controllers.GetCollection)
	http.HandleFunc("/collection/all", controllers.GetCollections)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
