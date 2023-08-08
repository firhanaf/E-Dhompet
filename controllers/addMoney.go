package controllers

import (
	"be18/acc-service/entities"
	"database/sql"
	"fmt"
	"log"
)

func AddMoney(db *sql.DB, loginUser entities.User) {
	users, loggedIn := Login(db, loginUser.Phone, loginUser.Password)
	for _, v := range users {
		if loggedIn {
			var topupAmount int64
			fmt.Print("Masukkan Jumlah Uang untuk Topup :")
			fmt.Scanln(&topupAmount)
			if topupAmount > 0 {
				v.Amount = topupAmount
				result, errTopup := db.Exec("insert into topup (user_id, amount, status) values (?, ?, ?)", v.Id, v.Amount, "SUCCESS")
				if errTopup != nil {
					log.Fatal("Error to Topup", errTopup.Error())
				} else {
					row, _ := result.RowsAffected()
					if row > 0 {
						fmt.Println("TOPUP SUCCESS")
						addition := v.Balance + v.Amount
						update, errUpdate := db.Exec("update users set balance = ? WHERE id = ?", addition, v.Id)
						if errUpdate != nil {
							log.Fatal("TOP UP FAILED", errUpdate.Error())
						} else {
							row, errRow := update.RowsAffected()
							if row > 0 {
								fmt.Println("SALDO BERTAMBAH")
							} else {
								log.Fatal(errRow.Error())
								fmt.Println("FAILED, SALDO BELUM BERTAMBAH")
							}
						}
					} else {
						fmt.Println("FAILED TO TOPUP")
					}
				}

			}

		}
	}
}
