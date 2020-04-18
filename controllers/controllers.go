package controllers

import (
	"log"
	"net/http"
	"strings"
)

func logerr(n int, err error) {
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

func CollectionID(w http.ResponseWriter, r *http.Request) (string, error) {
	param := strings.TrimPrefix(r.URL.Path, "/collection/")
	return param, nil
}
