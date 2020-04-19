package models

import (
	"cd_rental/db"
	"database/sql"
	"errors"
	"strconv"
	"time"
)

type User struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

type UserResponse struct {
	ID          int       `json:"id"`
	DateTime    time.Time `json:"dateTime"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phoneNumber"`
	Address     string    `json:"address"`
}

type UsersResponse struct {
	UsersResponse []UserResponse `json:"users"`
}

const (
	createUserQuery = "INSERT INTO users(date_time, name, phone_number, address) VALUES($1, $2, $3, $4) RETURNING id"
	getUserQuery    = "SELECT id, date_time, name, phone_number, address FROM users WHERE id = $1"
	getUsersQuery   = "SELECT id, date_time, name, phone_number, address FROM users ORDER BY id"
)

func NewUserResponse(id int, dateTime time.Time, name string, phoneNumber string, address string) UserResponse {
	newUserResponse := UserResponse{id, dateTime, name, phoneNumber, address}
	return newUserResponse
}

func PostUser(singleUser User) (UserResponse, error) {
	con := db.ConnectionDB()

	userResponse := UserResponse{}
	userResponse.DateTime = time.Now()

	err := con.QueryRow(createUserQuery,
		userResponse.DateTime,
		singleUser.Name,
		singleUser.PhoneNumber,
		singleUser.Address).Scan(&userResponse.ID)

	if err != nil {
		return userResponse, err
	}

	userResponse.Name = singleUser.Name
	userResponse.PhoneNumber = singleUser.PhoneNumber
	userResponse.Address = singleUser.Address

	return userResponse, nil
}

func GetUser(UserID string) (UserResponse, error) {
	con := db.ConnectionDB()

	userResponse := UserResponse{}

	userID, err := strconv.Atoi(UserID)
	if err != nil {
		return userResponse, err
	}

	err1 := con.QueryRow(getUserQuery, userID).
		Scan(&userResponse.ID,
			&userResponse.DateTime,
			&userResponse.Name,
			&userResponse.PhoneNumber,
			&userResponse.Address)

	if err1 == sql.ErrNoRows {
		return userResponse, errors.New("Queue is not found!")
	}

	if err1 != nil {
		return userResponse, err
	}

	return userResponse, nil
}

func GetUsers() (UsersResponse, error) {
	con := db.ConnectionDB()

	usersResponse := UsersResponse{}

	rows, err := con.Query(getUsersQuery)
	if err != nil {
		return usersResponse, err
	}

	for rows.Next() {
		userResponse := UserResponse{}

		err1 := rows.Scan(&userResponse.ID,
			&userResponse.DateTime,
			&userResponse.Name,
			&userResponse.PhoneNumber,
			&userResponse.Address)

		if err1 != nil {
			return usersResponse, err1
		}
		usersResponse.UsersResponse = append(usersResponse.UsersResponse, userResponse)
	}

	return usersResponse, nil
}
