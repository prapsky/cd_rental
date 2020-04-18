package controllers

import (
	"cd_rental/models"
	"encoding/json"
	"net/http"
)

func GetCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		collectionID, err := CollectionID(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		collection, err1 := models.GetCollection(collectionID)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}

		jsonInBytes, err2 := json.Marshal(collection)
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusInternalServerError)
			return
		}

		logerr(w.Write(jsonInBytes))
	default:
		w.WriteHeader(http.StatusNotFound)
		logerr(w.Write([]byte(`{"message": "not found"}`)))
	}
}

func PostCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		collection := models.Collection{}

		err := json.NewDecoder(r.Body).Decode(&collection)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err1 := models.PostCollection(collection)
		if err1 != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonInBytes, err2 := json.Marshal(response)
		if err2 != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		logerr(w.Write(jsonInBytes))
	default:
		w.WriteHeader(http.StatusNotFound)
		logerr(w.Write([]byte(`{"message": "not found"}`)))
	}
}
