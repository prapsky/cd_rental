package main

import (
	"cd_rental/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/collection", controllers.GetCollection)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
