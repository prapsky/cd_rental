package controllers

import (
	"cd_rental/models"
	"encoding/json"
	"net/http"
)

func PostReturnAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		returnRequest := models.ReturnRequest{}

		err := json.NewDecoder(r.Body).Decode(&returnRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err1 := models.PostReturnAll(returnRequest)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}

		jsonInBytes, err2 := json.Marshal(response)
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		logerr(w.Write(jsonInBytes))
	default:
		w.WriteHeader(http.StatusNotFound)
		logerr(w.Write([]byte(`{"message": "not found"}`)))
	}
}

func GetReturnAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		returnAllID, err := ReturnAllID(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		returnAll, err1 := models.GetReturnAll(returnAllID)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}

		jsonInBytes, err2 := json.Marshal(returnAll)
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
