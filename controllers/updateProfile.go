package controllers

import (
	"be18/acc-service/entities"
	"database/sql"
	"fmt"
	"log"
)

func UpdateUser(db *sql.DB, loginUser entities.User) {
	users, loggedIn := Login(db, loginUser.Phone, loginUser.Password)
	updateUser := entities.User{}
	for _, v := range users {
		if loggedIn {
			fmt.Println("Hello", v.Name)
			fmt.Println("Pilih Menu:\n1. Update email.\n2. Update password.\n3. Update address.\n4. Update Phone Number.\n0. Back to Menu")
			var menuUpdate int
			fmt.Println("input pilihan anda:")
			fmt.Scanln(&menuUpdate)
			if menuUpdate == 1 {
				fmt.Println("Input new Email:")
				fmt.Scanln(&updateUser.Email)
			}
			if menuUpdate == 2 {
				fmt.Println("Input new Password:")
				fmt.Scanln(&updateUser.Password)
			}
			if menuUpdate == 3 {
				fmt.Println("Input new Address:")
				fmt.Scanln(&updateUser.Address)
			}
			if menuUpdate == 4 {
				fmt.Println("Input new Phone Number:")
				fmt.Scanln(&updateUser.Phone)
			}
			if menuUpdate == 0 {
				return
			}

			if menuUpdate == 1 {
				result, errUpdate := db.Exec("UPDATE users SET email = ? WHERE phone = ?", updateUser.Email, loginUser.Phone)
				if errUpdate != nil {
					log.Fatal("error update email", errUpdate.Error())
				} else {
					row, _ := result.RowsAffected()
					if row > 0 {
						fmt.Println("EMAIL UPDATE SUCCESS")
					} else {
						fmt.Println("FAILED TO UPDATE EMAIL")
					}
				}
			}

			if menuUpdate == 2 {
				result, errUpdate2 := db.Exec("UPDATE users SET password = ? WHERE phone = ?", updateUser.Password, loginUser.Phone)
				if errUpdate2 != nil {
					log.Fatal("error update password", errUpdate2.Error())
				} else {
					row, _ := result.RowsAffected()
					if row > 0 {
						fmt.Println("PASSWORD UPDATE SUCCESS")
					} else {
						fmt.Println("FAILED TO UPDATE PASSWORD")
					}
				}
			}

			if menuUpdate == 3 {
				result, errUpdate3 := db.Exec("UPDATE users SET address = ? WHERE phone = ?", updateUser.Address, loginUser.Phone)
				if errUpdate3 != nil {
					log.Fatal("error update address", errUpdate3.Error())
				} else {
					row, _ := result.RowsAffected()
					if row > 0 {
						fmt.Println("ADDRESS UPDATE SUCCESSFULLY")
					} else {
						fmt.Println("FAILED TO UPDATE ADDRESS")
					}
				}
			}

			if menuUpdate == 4 {
				result, errUpdate4 := db.Exec("UPDATE users SET phone_number = ? WHERE phone = ?", updateUser.Phone, loginUser.Phone)
				if errUpdate4 != nil {
					log.Fatal("error update address", errUpdate4.Error())
				} else {
					row, _ := result.RowsAffected()
					if row > 0 {
						fmt.Println("PHONE NUMBER UPDATE SUCCESSFULLY")
					} else {
						fmt.Println("FAILED TO UPDATE PHONE NUMBER")
					}
				}
			}
		}
	}
}
