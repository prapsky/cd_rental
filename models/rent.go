package models

import (
	"cd_rental/db"
	"database/sql"
	"errors"
	"strconv"
	"time"
)

type Rent struct {
	QueueNumber  int `json:"queueNumber"`
	UserID       int `json:"userId"`
	CdID         int `json:"cdId"`
	RentQuantity int `json:"rentQuantity"`
}

type RentResponse struct {
	ID           int       `json:"id"`
	DateTime     time.Time `json:"dateTime"`
	QueueNumber  int       `json:"queueNumber"`
	UserID       int       `json:"userId"`
	CdID         int       `json:"cdId"`
	RentQuantity int       `json:"rentQuantity"`
}

type RentsResponse struct {
	RentsResponse []RentResponse `json:"rents"`
}

const (
	createRentQuery = "INSERT INTO rent(date_time, queue_number, user_id, cd_id, rent_quantity) VALUES($1, $2, $3, $4, $5) RETURNING id"
	getRentQuery    = "SELECT id, date_time, queue_number, user_id, cd_id, rent_quantity FROM rent WHERE id = $1"
)

func NewRentResponse(id int, dateTime time.Time, queueNumber int, userId int, cdId int, rentQuantity int) RentResponse {
	newRentResponse := RentResponse{id, dateTime, queueNumber, userId, cdId, rentQuantity}
	return newRentResponse
}

func PostRent(singleRent Rent) (RentResponse, error) {
	con := db.ConnectionDB()

	rentResponse := RentResponse{}
	rentResponse.DateTime = time.Now()

	err := con.QueryRow(createRentQuery,
		rentResponse.DateTime,
		singleRent.QueueNumber,
		singleRent.UserID,
		singleRent.CdID,
		singleRent.RentQuantity).Scan(&rentResponse.ID)

	if err != nil {
		return rentResponse, err
	}

	rentResponse.QueueNumber = singleRent.QueueNumber
	rentResponse.UserID = singleRent.UserID
	rentResponse.CdID = singleRent.CdID
	rentResponse.RentQuantity = singleRent.RentQuantity

	return rentResponse, nil
}

func GetRent(RentID string) (RentResponse, error) {
	con := db.ConnectionDB()

	rentResponse := RentResponse{}

	rentID, err := strconv.Atoi(RentID)
	if err != nil {
		return rentResponse, err
	}

	err1 := con.QueryRow(getRentQuery, rentID).
		Scan(&rentResponse.ID,
			&rentResponse.DateTime,
			&rentResponse.QueueNumber,
			&rentResponse.UserID,
			&rentResponse.CdID,
			&rentResponse.RentQuantity)

	if err1 == sql.ErrNoRows {
		return rentResponse, errors.New("Queue is not found!")
	}

	if err1 != nil {
		return rentResponse, err
	}

	return rentResponse, nil
}
