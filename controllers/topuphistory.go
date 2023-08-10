package controllers

import (
	"be18/acc-service/entities"
	"database/sql"
	"log"
)

func HistoryTopup(db *sql.DB, loginUser entities.User) []entities.User {
	users, loggedIn := Login(db, loginUser.Phone, loginUser.Password)
	var topuphistory []entities.User

	for _, v := range users {
		if loggedIn {
			history, errHistory := db.Query("select id, user_id, amount, status, transaction_time from topup where user_id = ?", v.Id)
			if errHistory != nil {
				log.Fatal("error to read topup history", errHistory.Error())
			}

			for history.Next() {
				var DataHistory entities.User
				errScan := history.Scan(&DataHistory.Topup_Id, &DataHistory.Id, &DataHistory.Amount, &DataHistory.Status, &DataHistory.Transaction_time_topup)
				if errScan != nil {
					log.Fatal("error to scan data", errScan.Error())
				}
				topuphistory = append(topuphistory, DataHistory)
			}
			// fmt.Println("isi :", topuphistory)
		}

	}
	return topuphistory
}
