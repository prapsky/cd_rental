package models

import (
	"time"
)

type CollectionResponse struct {
	DateTime time.Time `json:"dateTime"`
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Category string    `json:"category"`
	Quantity int       `json:"quantity"`
	Rate     int       `json:"rate"`
}

func NewCollectionResponse(dateTime time.Time, id int, title string, category string, quantity int, rate int) CollectionResponse {
	newCollectionResponse := CollectionResponse{dateTime, id, title, category, quantity, rate}
	return newCollectionResponse
}
