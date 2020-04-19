package models

import (
	"cd_rental/db"
	"database/sql"
	"errors"
	"strconv"
	"time"
)

type RentAllResponse struct {
	ID            int            `json:"id"`
	DateTime      time.Time      `json:"dateTime"`
	QueueNumber   int            `json:"queueNumber"`
	UserID        int            `json:"userId"`
	RentsResponse []RentResponse `json:"rents"`
}

const (
	createRentAllQuery = "INSERT INTO rentall(date_time, queue_number) VALUES($1, $2) RETURNING id"
	getRentAllQuery    = "SELECT id, date_time, queue_number FROM rentall WHERE id = $1"
	getRentsQuery      = "SELECT id, date_time, queue_number, user_id, cd_id, rent_quantity FROM rent ORDER BY id"
)

func PostRentAll(rentsResponse RentsResponse) (RentAllResponse, error) {
	con := db.ConnectionDB()

	rentAllResponse := RentAllResponse{}
	rentAllResponse.DateTime = time.Now()

	err := con.QueryRow(createRentAllQuery,
		rentAllResponse.DateTime,
		rentsResponse.RentsResponse[0].QueueNumber).Scan(&rentAllResponse.ID)

	if err != nil {
		return rentAllResponse, err
	}

	rentAllResponse.QueueNumber = rentsResponse.RentsResponse[0].QueueNumber
	rentAllResponse.UserID = rentsResponse.RentsResponse[0].UserID
	rentAllResponse.RentsResponse = rentsResponse.RentsResponse

	return rentAllResponse, nil
}

func GetRentAll(RentAllID string) (RentAllResponse, error) {
	con := db.ConnectionDB()

	rentAllResponse := RentAllResponse{}

	rentAllID, err := strconv.Atoi(RentAllID)
	if err != nil {
		return rentAllResponse, err
	}

	err1 := con.QueryRow(getRentAllQuery, rentAllID).
		Scan(&rentAllResponse.ID,
			&rentAllResponse.DateTime,
			&rentAllResponse.QueueNumber)

	if err1 == sql.ErrNoRows {
		return rentAllResponse, errors.New("Queue is not found!")
	}

	if err1 != nil {
		return rentAllResponse, err
	}

	rentsResponse := RentsResponse{}

	rows, err2 := con.Query(getRentsQuery)
	if err2 != nil {
		return rentAllResponse, err2
	}

	for rows.Next() {
		rentResponse := RentResponse{}

		err3 := rows.Scan(&rentResponse.ID,
			&rentResponse.DateTime,
			&rentResponse.QueueNumber,
			&rentResponse.UserID,
			&rentResponse.CdID,
			&rentResponse.RentQuantity)

		if err3 != nil {
			return rentAllResponse, err3
		}
		rentsResponse.RentsResponse = append(rentsResponse.RentsResponse, rentResponse)
	}

	rentAllResponse.UserID = rentsResponse.RentsResponse[0].UserID
	rentAllResponse.RentsResponse = rentsResponse.RentsResponse

	return rentAllResponse, nil
}
