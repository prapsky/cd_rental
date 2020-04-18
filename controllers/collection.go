package controllers

import (
	"cd_rental/models"
	"encoding/json"
	"net/http"
	"time"
)

func GetCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		collection := models.NewCollectionResponse(
			time.Date(2019, 01, 01, 16, 00, 0, 0, time.UTC),
			1,
			"Star Wars",
			"Sci-Fi",
			20,
			15000,
		)

		jsonInBytes, err := json.Marshal(collection)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logerr(w.Write(jsonInBytes))
	default:
		w.WriteHeader(http.StatusNotFound)
		logerr(w.Write([]byte(`{"message": "not found"}`)))
	}
}
