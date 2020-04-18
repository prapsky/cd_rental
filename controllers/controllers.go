package controllers

import (
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

func logerr(n int, err error) {
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

func CollectionID(w http.ResponseWriter, r *http.Request) (string, error) {
	param := strings.TrimPrefix(r.URL.Path, "/collection/")
	index := strings.Index(param, "/")
	static := param[index+1:]

	if len(static) == 0 {
		http.Redirect(w, r, r.URL.Path[:len(r.URL.Path)-1], http.StatusMovedPermanently)
		return "error!", nil
	}

	param = path.Clean(param)
	collectionID := filepath.Dir(param)

	return collectionID, nil
}
