package db

import (
	"log"

	"database/sql"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func ConnectionDB() *sql.DB {
	var err error
	DB, err = sql.Open("postgres", "user=prapsky dbname=cdrental sslmode=disable")
	if err != nil {
		log.Println(err.Error())
	}

	err = DB.Ping()
	log.Println(err)
	if err != nil {
		log.Println("DB IS NOT connected")
		log.Println(err.Error())
	}

	return DB
}
