package controllers

import (
	"cd_rental/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func PostRent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		rentRequest := models.RentRequest{}

		err := json.NewDecoder(r.Body).Decode(&rentRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err1 := models.PostRent(rentRequest)
		if err1 != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		collection, err2 := models.GetCollection(strconv.Itoa(response.CdID))
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusInternalServerError)
			return
		}

		updateQuantityCollection := models.UpdateQuantityCollection{}
		updateQuantityCollection.ID = response.CdID
		updateQuantityCollection.Quantity = collection.Quantity - response.RentQuantity

		_, err3 := models.PatchCollection(updateQuantityCollection)
		if err3 != nil {
			http.Error(w, err3.Error(), http.StatusInternalServerError)
			return
		}

		jsonInBytes, err4 := json.Marshal(response)
		if err4 != nil {
			http.Error(w, err4.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		logerr(w.Write(jsonInBytes))

	default:
		w.WriteHeader(http.StatusNotFound)
		logerr(w.Write([]byte(`{"message": "not found"}`)))
	}
}

func Rent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		rentID, err := RentID(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rent, err1 := models.GetRent(rentID)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}

		jsonInBytes, err2 := json.Marshal(rent)
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

func RentQueue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		queueNumber, err := QueueNumber(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rentQueue, err1 := models.GetRentQueue(queueNumber)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}

		jsonInBytes, err2 := json.Marshal(rentQueue)
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
