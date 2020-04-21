package models

import (
	"cd_rental/db"
	"database/sql"
	"errors"
	"math"
	"strconv"
	"time"
)

type PaymentRequest struct {
	ReturnID        int              `json:"returnId"`
	UserID          int              `json:"userId"`
	ReturnsResponse []ReturnResponse `json:"returns"`
}

type PaymentResponse struct {
	ID              int              `json:"id"`
	DateTime        time.Time        `json:"dateTime"`
	ReturnID        int              `json:"returnId"`
	UserID          int              `json:"userId"`
	TotalPayment    int              `json:"totalPayment"`
	ReturnsResponse []ReturnResponse `json:"returns"`
}

const (
	createPaymentQuery = "INSERT INTO payment(date_time, return_id, total_payment) VALUES($1, $2, $3) RETURNING id"
	getPaymentQuery    = "SELECT id, date_time, return_id, total_payment FROM payment WHERE id = $1"
)

func PostPayment(paymentRequest PaymentRequest) (PaymentResponse, error) {
	con := db.ConnectionDB()

	paymentResponse := PaymentResponse{}
	paymentResponse.ReturnID = paymentRequest.ReturnID
	paymentResponse.UserID = paymentRequest.UserID
	paymentResponse.DateTime = time.Now()

	returnAllResponse := ReturnAllResponse{}
	err := con.QueryRow(getReturnAllQuery, paymentResponse.ReturnID).
		Scan(&returnAllResponse.ID,
			&returnAllResponse.DateTime,
			&returnAllResponse.RentAllID)
	if err == sql.ErrNoRows {
		return paymentResponse, errors.New("Queue is not found!")
	}
	if err != nil {
		return paymentResponse, err
	}

	rentAllResponse := RentAllResponse{}
	err1 := con.QueryRow(getRentAllQuery, returnAllResponse.RentAllID).
		Scan(&rentAllResponse.ID,
			&rentAllResponse.DateTime,
			&rentAllResponse.QueueNumber)
	if err1 == sql.ErrNoRows {
		return paymentResponse, errors.New("Queue is not found!")
	}
	if err1 != nil {
		return paymentResponse, err1
	}

	rentsResponse := RentsResponse{}
	rows, err2 := con.Query(getRentsQuery, rentAllResponse.QueueNumber)
	if err2 != nil {
		return paymentResponse, err2
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
			return paymentResponse, err3
		}
		rentsResponse.RentsResponse = append(rentsResponse.RentsResponse, rentResponse)
	}

	returnsResponse := ReturnsResponse{}
	lengthRentsResponse := len(paymentRequest.ReturnsResponse)
	returnsResponse.ReturnsResponse = make([]ReturnResponse, lengthRentsResponse)
	for i := 0; i < lengthRentsResponse; i++ {
		returnsResponse.ReturnsResponse[i].CdID = rentsResponse.RentsResponse[i].CdID
		returnsResponse.ReturnsResponse[i].ReturnQuantity = rentsResponse.RentsResponse[i].RentQuantity

		difference := (paymentResponse.DateTime).Sub(rentsResponse.RentsResponse[i].DateTime)
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
			return paymentResponse, errors.New("Queue is not found!")
		}
		if err4 != nil {
			return paymentResponse, err4
		}
		returnsResponse.ReturnsResponse[i].RatePerDay = collectionResponse.Rate
		returnsResponse.ReturnsResponse[i].TotalRate = returnsResponse.ReturnsResponse[i].RatePerDay * returnsResponse.ReturnsResponse[i].RentDays * returnsResponse.ReturnsResponse[i].ReturnQuantity
		paymentResponse.TotalPayment = paymentResponse.TotalPayment + returnsResponse.ReturnsResponse[i].TotalRate
	}
	paymentResponse.ReturnsResponse = returnsResponse.ReturnsResponse

	err5 := con.QueryRow(createPaymentQuery,
		paymentResponse.DateTime,
		paymentRequest.ReturnID,
		paymentResponse.TotalPayment).Scan(&paymentResponse.ID)
	if err5 != nil {
		return paymentResponse, err5
	}

	return paymentResponse, nil
}

func GetPayment(PaymentID string) (PaymentResponse, error) {
	con := db.ConnectionDB()

	paymentResponse := PaymentResponse{}

	paymentID, err := strconv.Atoi(PaymentID)
	if err != nil {
		return paymentResponse, err
	}

	err1 := con.QueryRow(getPaymentQuery, paymentID).
		Scan(&paymentResponse.ID,
			&paymentResponse.DateTime,
			&paymentResponse.ReturnID,
			&paymentResponse.TotalPayment)

	if err1 == sql.ErrNoRows {
		return paymentResponse, errors.New("Queue is not found!")
	}

	if err1 != nil {
		return paymentResponse, err1
	}

	returnAllResponse := ReturnAllResponse{}
	err2 := con.QueryRow(getReturnAllQuery, paymentResponse.ReturnID).
		Scan(&returnAllResponse.ID,
			&returnAllResponse.DateTime,
			&returnAllResponse.RentAllID)
	if err2 == sql.ErrNoRows {
		return paymentResponse, errors.New("Queue is not found!")
	}
	if err2 != nil {
		return paymentResponse, err2
	}

	rentAllResponse := RentAllResponse{}
	err3 := con.QueryRow(getRentAllQuery, returnAllResponse.RentAllID).
		Scan(&rentAllResponse.ID,
			&rentAllResponse.DateTime,
			&rentAllResponse.QueueNumber)
	if err3 == sql.ErrNoRows {
		return paymentResponse, errors.New("Queue is not found!")
	}
	if err3 != nil {
		return paymentResponse, err3
	}

	rentsResponse := RentsResponse{}
	rows, err4 := con.Query(getRentsQuery, rentAllResponse.QueueNumber)
	if err4 != nil {
		return paymentResponse, err4
	}
	for rows.Next() {
		rentResponse := RentResponse{}
		err5 := rows.Scan(&rentResponse.ID,
			&rentResponse.DateTime,
			&rentResponse.QueueNumber,
			&rentResponse.UserID,
			&rentResponse.CdID,
			&rentResponse.RentQuantity)
		if err5 != nil {
			return paymentResponse, err5
		}
		rentsResponse.RentsResponse = append(rentsResponse.RentsResponse, rentResponse)
	}

	returnsResponse := ReturnsResponse{}
	lengthRentsResponse := len(rentsResponse.RentsResponse)
	returnsResponse.ReturnsResponse = make([]ReturnResponse, lengthRentsResponse)
	for i := 0; i < lengthRentsResponse; i++ {
		returnsResponse.ReturnsResponse[i].CdID = rentsResponse.RentsResponse[i].CdID
		returnsResponse.ReturnsResponse[i].ReturnQuantity = rentsResponse.RentsResponse[i].RentQuantity

		difference := (paymentResponse.DateTime).Sub(rentsResponse.RentsResponse[i].DateTime)
		difference_hours := difference.Hours()
		twentyfourhours := 24.0
		days := math.Round(difference_hours / twentyfourhours)
		returnsResponse.ReturnsResponse[i].RentDays = int(days)

		collectionResponse := CollectionResponse{}

		err6 := con.QueryRow(getCollectionQuery, returnsResponse.ReturnsResponse[i].CdID).
			Scan(&collectionResponse.ID,
				&collectionResponse.DateTime,
				&collectionResponse.Title,
				&collectionResponse.Category,
				&collectionResponse.Quantity,
				&collectionResponse.Rate)
		if err6 == sql.ErrNoRows {
			return paymentResponse, errors.New("Queue is not found!")
		}
		if err6 != nil {
			return paymentResponse, err6
		}
		returnsResponse.ReturnsResponse[i].RatePerDay = collectionResponse.Rate
		returnsResponse.ReturnsResponse[i].TotalRate = returnsResponse.ReturnsResponse[i].RatePerDay * returnsResponse.ReturnsResponse[i].RentDays * returnsResponse.ReturnsResponse[i].ReturnQuantity
	}
	paymentResponse.ReturnsResponse = returnsResponse.ReturnsResponse

	paymentResponse.UserID = rentsResponse.RentsResponse[0].UserID

	return paymentResponse, nil
}
