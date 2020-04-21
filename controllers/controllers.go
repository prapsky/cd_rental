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

func UserID(w http.ResponseWriter, r *http.Request) (string, error) {
	param := strings.TrimPrefix(r.URL.Path, "/user/")
	return param, nil
}

func RentID(w http.ResponseWriter, r *http.Request) (string, error) {
	param := strings.TrimPrefix(r.URL.Path, "/rent/")
	return param, nil
}

func QueueNumber(w http.ResponseWriter, r *http.Request) (string, error) {
	param := strings.TrimPrefix(r.URL.Path, "/rent/queue/")
	return param, nil
}

func RentAllID(w http.ResponseWriter, r *http.Request) (string, error) {
	param := strings.TrimPrefix(r.URL.Path, "/rent/all/")
	return param, nil
}

func ReturnAllID(w http.ResponseWriter, r *http.Request) (string, error) {
	param := strings.TrimPrefix(r.URL.Path, "/return/all/")
	return param, nil
}

func PaymentID(w http.ResponseWriter, r *http.Request) (string, error) {
	param := strings.TrimPrefix(r.URL.Path, "/payment/")
	return param, nil
}
