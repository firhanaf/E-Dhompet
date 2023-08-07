package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func GetDB() (*sql.DB, error) {

	var db *sql.DB
	var err error

	//Setting db connection
	var connectionString = os.Getenv("CONNECTION_DB")
	// fmt.Println("connectionstring:", connectionString)
	db, err = sql.Open("mysql", connectionString)
	if err != nil { // jika terjadi error
		log.Fatal("error open connection to db ", err.Error())
	}
	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("error connect to db ", errPing.Error())
	} else {
		fmt.Println("CONNECTED")
	}
	return db, nil
}
