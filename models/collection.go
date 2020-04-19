package models

import (
	"cd_rental/db"
	"database/sql"
	"errors"
	"strconv"
	"time"
)

type Collection struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Quantity int    `json:"quantity"`
	Rate     int    `json:"rate"`
}

type UpdateCollection struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Quantity int    `json:"quantity"`
	Rate     int    `json:"rate"`
}

type UpdateQuantityCollection struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}

type CollectionResponse struct {
	ID       int       `json:"id"`
	DateTime time.Time `json:"dateTime"`
	Title    string    `json:"title"`
	Category string    `json:"category"`
	Quantity int       `json:"quantity"`
	Rate     int       `json:"rate"`
}

type CollectionsResponse struct {
	CollectionsResponse []CollectionResponse `json:"collections"`
}

const (
	createCollectionQuery         = "INSERT INTO collection(date_time, title, category, quantity, rate) VALUES($1, $2, $3, $4, $5) RETURNING id"
	getCollectionQuery            = "SELECT id, date_time, title, category, quantity, rate FROM collection WHERE id = $1"
	getCollectionsQuery           = "SELECT id, date_time, title, category, quantity, rate FROM collection ORDER BY id"
	updateCollectionQuery         = "UPDATE collection SET date_time=$1, title=$2, category=$3, quantity=$4, rate=$5 WHERE id=$6"
	updateQuantityCollectionQuery = "UPDATE collection SET date_time=$1, quantity=$2 WHERE id=$3"
)

func NewCollectionResponse(id int, dateTime time.Time, title string, category string, quantity int, rate int) CollectionResponse {
	newCollectionResponse := CollectionResponse{id, dateTime, title, category, quantity, rate}
	return newCollectionResponse
}

func PostCollection(singleCollection Collection) (CollectionResponse, error) {
	con := db.ConnectionDB()

	collectionResponse := CollectionResponse{}
	collectionResponse.DateTime = time.Now()

	err := con.QueryRow(createCollectionQuery,
		collectionResponse.DateTime,
		singleCollection.Title,
		singleCollection.Category,
		singleCollection.Quantity,
		singleCollection.Rate).Scan(&collectionResponse.ID)

	if err != nil {
		return collectionResponse, err
	}

	collectionResponse.Title = singleCollection.Title
	collectionResponse.Category = singleCollection.Category
	collectionResponse.Quantity = singleCollection.Quantity
	collectionResponse.Rate = singleCollection.Rate

	return collectionResponse, nil
}

func GetCollection(CollectionID string) (CollectionResponse, error) {
	con := db.ConnectionDB()

	collectionResponse := CollectionResponse{}

	cdID, err := strconv.Atoi(CollectionID)
	if err != nil {
		return collectionResponse, err
	}

	err1 := con.QueryRow(getCollectionQuery, cdID).
		Scan(&collectionResponse.ID,
			&collectionResponse.DateTime,
			&collectionResponse.Title,
			&collectionResponse.Category,
			&collectionResponse.Quantity,
			&collectionResponse.Rate)

	if err1 == sql.ErrNoRows {
		return collectionResponse, errors.New("Queue is not found!")
	}

	if err1 != nil {
		return collectionResponse, err
	}

	return collectionResponse, nil
}

func GetCollections() (CollectionsResponse, error) {
	con := db.ConnectionDB()

	collectionsResponse := CollectionsResponse{}

	rows, err := con.Query(getCollectionsQuery)
	if err != nil {
		return collectionsResponse, err
	}

	for rows.Next() {
		collectionResponse := CollectionResponse{}

		err2 := rows.Scan(&collectionResponse.ID,
			&collectionResponse.DateTime,
			&collectionResponse.Title,
			&collectionResponse.Category,
			&collectionResponse.Quantity,
			&collectionResponse.Rate)

		if err2 != nil {
			return collectionsResponse, err2
		}
		collectionsResponse.CollectionsResponse = append(collectionsResponse.CollectionsResponse, collectionResponse)
	}

	return collectionsResponse, nil
}

func PutCollection(singleUpdateCollection UpdateCollection) (CollectionResponse, error) {
	con := db.ConnectionDB()

	collectionResponse := CollectionResponse{}

	collectionResponse.DateTime = time.Now()

	_, err := con.Exec(updateCollectionQuery,
		collectionResponse.DateTime,
		singleUpdateCollection.Title,
		singleUpdateCollection.Category,
		singleUpdateCollection.Quantity,
		singleUpdateCollection.Rate,
		singleUpdateCollection.ID)

	if err != nil {
		return collectionResponse, err
	}

	collectionResponse.ID = singleUpdateCollection.ID
	collectionResponse.Title = singleUpdateCollection.Title
	collectionResponse.Category = singleUpdateCollection.Category
	collectionResponse.Quantity = singleUpdateCollection.Quantity
	collectionResponse.Rate = singleUpdateCollection.Rate

	return collectionResponse, nil
}

func PatchCollection(updateQuantityCollection UpdateQuantityCollection) (CollectionResponse, error) {
	con := db.ConnectionDB()

	collectionResponse := CollectionResponse{}

	collectionResponse.DateTime = time.Now()

	_, err := con.Exec(updateQuantityCollectionQuery,
		collectionResponse.DateTime,
		updateQuantityCollection.Quantity,
		updateQuantityCollection.ID)

	if err != nil {
		return collectionResponse, err
	}

	err1 := con.QueryRow(getCollectionQuery, updateQuantityCollection.ID).
		Scan(&collectionResponse.ID,
			&collectionResponse.DateTime,
			&collectionResponse.Title,
			&collectionResponse.Category,
			&collectionResponse.Quantity,
			&collectionResponse.Rate)

	if err1 == sql.ErrNoRows {
		return collectionResponse, errors.New("Queue is not found!")
	}

	if err1 != nil {
		return collectionResponse, err
	}

	return collectionResponse, nil
}
