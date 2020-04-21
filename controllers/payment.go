package controllers

import (
	"cd_rental/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func PostPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		paymentRequest := models.PaymentRequest{}

		err := json.NewDecoder(r.Body).Decode(&paymentRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err1 := models.PostPayment(paymentRequest)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}

		for i := 0; i < len(paymentRequest.ReturnsResponse); i++ {
			collection, err2 := models.GetCollection(strconv.Itoa(response.ReturnsResponse[i].CdID))
			if err2 != nil {
				http.Error(w, err2.Error(), http.StatusInternalServerError)
				return
			}

			updateQuantityCollection := models.UpdateQuantityCollection{}
			updateQuantityCollection.ID = response.ReturnsResponse[i].CdID
			updateQuantityCollection.Quantity = collection.Quantity + response.ReturnsResponse[i].ReturnQuantity

			_, err3 := models.PatchCollection(updateQuantityCollection)
			if err3 != nil {
				http.Error(w, err3.Error(), http.StatusInternalServerError)
				return
			}
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

func GetPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		paymentID, err := PaymentID(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		payment, err1 := models.GetPayment(paymentID)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}

		jsonInBytes, err2 := json.Marshal(payment)
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
