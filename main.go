package main

import (
	"cd_rental/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/collection", controllers.PostCollection)
	http.HandleFunc("/collection/", controllers.Collection)
	http.HandleFunc("/collection/all", controllers.GetCollections)

	http.HandleFunc("/user", controllers.PostUser)
	http.HandleFunc("/user/", controllers.GetUser)
	http.HandleFunc("/user/all", controllers.GetUsers)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
