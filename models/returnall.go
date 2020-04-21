package models

import (
	"cd_rental/db"
	"database/sql"
	"errors"
	"math"
	"strconv"
	"time"
)

type ReturnRequest struct {
	RentAllID     int            `json:"rentAllId"`
	QueueNumber   int            `json:"queueNumber"`
	UserID        int            `json:"userId"`
	RentsResponse []RentResponse `json:"rents"`
}

type ReturnResponse struct {
	CdID           int `json:"cdId"`
	ReturnQuantity int `json:"returnQuantity"`
	RentDays       int `json:"rentDays"`
	RatePerDay     int `json:"ratePerDay"`
	TotalRate      int `json:"totalRate"`
}

type ReturnsResponse struct {
	ReturnsResponse []ReturnResponse `json:"returns"`
}

type ReturnAllResponse struct {
	ID              int              `json:"id"`
	DateTime        time.Time        `json:"dateTime"`
	RentAllID       int              `json:"rentAllId"`
	UserID          int              `json:"userId"`
	ReturnsResponse []ReturnResponse `json:"returns"`
}

const (
	createReturnAllQuery = "INSERT INTO return(date_time, rent_all_id) VALUES($1, $2) RETURNING id"
	getReturnAllQuery    = "SELECT id, date_time, rent_all_id FROM return WHERE id = $1"
)

func NewReturnResponse(cdId int, returnQuantity int, rentDays int, ratePerDay int, totalRate int) ReturnResponse {
	newReturnResponse := ReturnResponse{cdId, returnQuantity, rentDays, ratePerDay, totalRate}
	return newReturnResponse
}

func PostReturnAll(returnRequest ReturnRequest) (ReturnAllResponse, error) {
	con := db.ConnectionDB()

	returnAllResponse := ReturnAllResponse{}

	returnAllResponse.DateTime = time.Now()

	err := con.QueryRow(createReturnAllQuery,
		returnAllResponse.DateTime,
		returnRequest.RentAllID).Scan(&returnAllResponse.ID)
	if err != nil {
		return returnAllResponse, err
	}

	returnAllResponse.RentAllID = returnRequest.RentAllID

	returnAllResponse.UserID = returnRequest.UserID

	rentsResponse := RentsResponse{}
	rows, err2 := con.Query(getRentsQuery, returnRequest.QueueNumber)
	if err2 != nil {
		return returnAllResponse, err2
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
			return returnAllResponse, err3
		}
		rentsResponse.RentsResponse = append(rentsResponse.RentsResponse, rentResponse)
	}

	returnsResponse := ReturnsResponse{}
	lengthRentsResponse := len(returnRequest.RentsResponse)
	returnsResponse.ReturnsResponse = make([]ReturnResponse, lengthRentsResponse)

	for i := 0; i < lengthRentsResponse; i++ {
		returnsResponse.ReturnsResponse[i].CdID = rentsResponse.RentsResponse[i].CdID
		returnsResponse.ReturnsResponse[i].ReturnQuantity = rentsResponse.RentsResponse[i].RentQuantity

		difference := (returnAllResponse.DateTime).Sub(returnRequest.RentsResponse[i].DateTime)
		difference_hours := difference.Hours()
		twentyfourhours := 24.0
		days := math.Round(difference_hours / twentyfourhours)
		returnsResponse.ReturnsResponse[i].RentDays = int(days)

		collectionResponse := CollectionResponse{}

		err4 := con.QueryRow(getCollectionQuery, returnsResponse.ReturnsResponse[i].CdID).
			Scan(&collectionResponse.ID,
				&collectionResponse.DateTime,
				&collectionResponse.Title,
				&collectionResponse.Category,
				&collectionResponse.Quantity,
				&collectionResponse.Rate)
		if err4 == sql.ErrNoRows {
			return returnAllResponse, errors.New("Queue is not found!")
		}
		if err4 != nil {
			return returnAllResponse, err
		}
		returnsResponse.ReturnsResponse[i].RatePerDay = collectionResponse.Rate
		returnsResponse.ReturnsResponse[i].TotalRate = returnsResponse.ReturnsResponse[i].RatePerDay * returnsResponse.ReturnsResponse[i].RentDays * returnsResponse.ReturnsResponse[i].ReturnQuantity
	}
	returnAllResponse.ReturnsResponse = returnsResponse.ReturnsResponse

	return returnAllResponse, nil
}

func GetReturnAll(ReturnAllID string) (ReturnAllResponse, error) {
	con := db.ConnectionDB()

	returnAllResponse := ReturnAllResponse{}
	returnAllID, err := strconv.Atoi(ReturnAllID)
	if err != nil {
		return returnAllResponse, err
	}
	err1 := con.QueryRow(getReturnAllQuery, returnAllID).
		Scan(&returnAllResponse.ID,
			&returnAllResponse.DateTime,
			&returnAllResponse.RentAllID)
	if err1 == sql.ErrNoRows {
		return returnAllResponse, errors.New("Queue is not found!")
	}
	if err1 != nil {
		return returnAllResponse, err
	}

	rentAllResponse := RentAllResponse{}
	err2 := con.QueryRow(getRentAllQuery, returnAllResponse.RentAllID).
		Scan(&rentAllResponse.ID,
			&rentAllResponse.DateTime,
			&rentAllResponse.QueueNumber)
	if err2 == sql.ErrNoRows {
		return returnAllResponse, errors.New("Queue is not found!")
	}
	if err2 != nil {
		return returnAllResponse, err
	}

	rentsResponse := RentsResponse{}
	rows, err3 := con.Query(getRentsQuery, rentAllResponse.QueueNumber)
	if err3 != nil {
		return returnAllResponse, err3
	}
	for rows.Next() {
		rentResponse := RentResponse{}
		err4 := rows.Scan(&rentResponse.ID,
			&rentResponse.DateTime,
			&rentResponse.QueueNumber,
			&rentResponse.UserID,
			&rentResponse.CdID,
			&rentResponse.RentQuantity)
		if err4 != nil {
			return returnAllResponse, err4
		}
		rentsResponse.RentsResponse = append(rentsResponse.RentsResponse, rentResponse)
	}
	returnsResponse := ReturnsResponse{}
	lengthRentsResponse := len(rentsResponse.RentsResponse)
	returnsResponse.ReturnsResponse = make([]ReturnResponse, lengthRentsResponse)
	for i := 0; i < lengthRentsResponse; i++ {
		returnsResponse.ReturnsResponse[i].CdID = rentsResponse.RentsResponse[i].CdID

		returnsResponse.ReturnsResponse[i].ReturnQuantity = rentsResponse.RentsResponse[i].RentQuantity

		difference := (returnAllResponse.DateTime).Sub(rentsResponse.RentsResponse[i].DateTime)
		difference_hours := difference.Hours()
		twentyfourhours := 24.0
		days := math.Round(difference_hours / twentyfourhours)
		returnsResponse.ReturnsResponse[i].RentDays = int(days)

		collectionResponse := CollectionResponse{}
		err5 := con.QueryRow(getCollectionQuery, returnsResponse.ReturnsResponse[i].CdID).
			Scan(&collectionResponse.ID,
				&collectionResponse.DateTime,
				&collectionResponse.Title,
				&collectionResponse.Category,
				&collectionResponse.Quantity,
				&collectionResponse.Rate)
		if err5 == sql.ErrNoRows {
			return returnAllResponse, errors.New("Queue is not found!")
		}
		if err5 != nil {
			return returnAllResponse, err
		}
		returnsResponse.ReturnsResponse[i].RatePerDay = collectionResponse.Rate
		returnsResponse.ReturnsResponse[i].TotalRate = returnsResponse.ReturnsResponse[i].RatePerDay * returnsResponse.ReturnsResponse[i].RentDays * returnsResponse.ReturnsResponse[i].ReturnQuantity
	}
	returnAllResponse.ReturnsResponse = returnsResponse.ReturnsResponse
	returnAllResponse.UserID = rentsResponse.RentsResponse[0].UserID

	return returnAllResponse, nil
}
