package controllers

import (
	"be18/acc-service/entities"
	"database/sql"
	"fmt"
)

func ReadProfile(db *sql.DB, loginUser entities.User) {
	users, loggedIn := Login(db, loginUser.Phone, loginUser.Password)
	for _, v := range users {
		if loggedIn {
			fmt.Println("ID : ", v.Id)
			fmt.Println("Name : ", v.Name)
			fmt.Println("Username : ", v.Username)
			fmt.Println("Phone : ", v.Phone)
			fmt.Println("Address : ", v.Address)
			fmt.Println("Balance : ", v.Balance)
			fmt.Println("============================\n")
		} else {
			fmt.Println("Access Denied")
		}
	}
}
