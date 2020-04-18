package main

import (
	"fmt"
	"log"
	"net/http"
)

func getCollection(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GET /collection")
}

func main() {
	http.HandleFunc("/collection", getCollection)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
