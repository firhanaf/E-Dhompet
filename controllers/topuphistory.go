package controllers

import (
	"be18/acc-service/entities"
	"database/sql"
	"log"
)

func HistoryTopup(db *sql.DB, loginUser entities.User) []entities.User {
	users, loggedIn := Login(db, loginUser.Phone, loginUser.Password)
	var topuphistory []entities.User

	query := "select users.id, users.name, topup.id, topup.amount, topup.status, users.balance, topup.transaction_time from users inner join topup on users.id = topup.user_id where user_id = ? ORDER BY transaction_time desc"

	for _, v := range users {
		if loggedIn {
			history, errHistory := db.Query(query, v.Id)
			if errHistory != nil {
				log.Fatal("error to read topup history", errHistory.Error())
			}

			for history.Next() {
				var DataHistory entities.User
				errScan := history.Scan(&DataHistory.Id, &DataHistory.Name, &DataHistory.Topup_Id, &DataHistory.Amount, &DataHistory.Status, &DataHistory.Balance, &DataHistory.Transaction_time_topup)
				if errScan != nil {
					log.Fatal("error to scan data", errScan.Error())
				}
				DataHistory.Balance -= DataHistory.Amount
				topuphistory = append(topuphistory, DataHistory)
			}
			// fmt.Println("isi :", topuphistory)
		}

	}
	return topuphistory
}
