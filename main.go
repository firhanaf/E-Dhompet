package main

import (
	"fmt"
	"log"

	"be18/acc-service/controllers"
	"be18/acc-service/entities"

	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := controllers.GetDB()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}
	defer db.Close()

	//Main Menu
	var loggedInUser entities.User
	for {
		fmt.Println("Selamat Datang di E-Dhompet")
		fmt.Println("Pilih Menu: \n 1. Register \n 2. Login \n 3. View Profile \n 4. Update Profile \n 5. Delete Account \n 6. Top Up \n 7. Transfer \n 8. Top Up History \n 9. Transfer History \n 10. View Profile by Phone Number \n 0. Log Out")
		var menu int
		fmt.Println("input pilihan anda:")
		fmt.Scanln(&menu)

		switch menu {
		case 1: // fitur Register

			// biasanya urutannya email, username, password, nama lengkap, nomor telepon, address

			registerUser := entities.User{}
			// fmt.Println("Masukkan ID :") // Email
			// fmt.Scanln(&registerUser.Id)
			fmt.Println("Masukkan Email :") // Email
			fmt.Scanln(&registerUser.Email)
			fmt.Println("Masukkan Username :") // Username
			fmt.Scanln(&registerUser.Username)
			fmt.Println("Masukkan Password :") // Password
			fmt.Scanln(&registerUser.Password)
			fmt.Println("Masukkan Nama :") // Nama Lengkap
			fmt.Scanln(&registerUser.Name)
			fmt.Println("Masukkan Nomor Telepon :") // Nomor Telepon
			fmt.Scanln(&registerUser.Phone)
			fmt.Println("Masukkan Adress :") // Address
			fmt.Scanln(&registerUser.Address)

			uuid := uuid.New()

			registerRows, errRegister := db.Exec("INSERT INTO Users (id, name, username, email, phone, password, address, balance) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", uuid, registerUser.Name, registerUser.Username, registerUser.Email, registerUser.Phone, registerUser.Password, registerUser.Address, 0)
			if errRegister != nil {
				log.Fatal("error insert", errRegister.Error())
			} else {
				row, _ := registerRows.RowsAffected()
				if row > 0 {
					fmt.Println("Success Insert Data")
				} else {
					fmt.Println("Failed to Insert Data")
				}
			}

		case 2: // fitur Login

			loginUser := entities.User{}
			fmt.Println("Masukkan Nomor Telepon :")
			fmt.Scanln(&loginUser.Phone)
			fmt.Println("Masukkan Password :")
			fmt.Scanln(&loginUser.Password)

			users, loggedIn := controllers.Login(db, loginUser.Phone, loginUser.Password)
			if loggedIn && len(users) > 0 {
				loggedInUser = users[0] // Store the logged-in user
				fmt.Println("Login Success!\nWelcome", loggedInUser.Name)
			} else {
				fmt.Println("Login Failed!")
			}

		case 3: // view profile yang telah login
			fmt.Println("View Profile")
			// loginUser := entities.User{}
			// fmt.Print("Masukkan Nomor Telepon : ")
			// fmt.Scanln(&loginUser.Phone)
			// fmt.Print("Masukkan Password : ")
			// fmt.Scanln(&loginUser.Password)

			// users, loggedin := controllers.Login(db, loginUser.Phone, loginUser.Password)
			// for _, v := range users {
			if loggedInUser.Id != "" {
				fmt.Println("=========================")
				fmt.Println("Displaying Profile of:", loggedInUser.Name)
				fmt.Println("=========================")
				controllers.ReadProfile(db, loggedInUser)
			} else {
				fmt.Println("Please login first!")
			}

		case 4: // update profil

			fmt.Println("Update Profile")
			// loginUser := entities.User{}
			// fmt.Print("Masukkan Nomor Telepon : ")
			// fmt.Scanln(&loginUser.Phone)
			// fmt.Print("Masukkan Password : ")
			// fmt.Scanln(&loginUser.Password)
			// users, loggedin := controllers.Login(db, loginUser.Phone, loginUser.Password)
			// for _, v := range users {
			if loggedInUser.Id != "" {
				controllers.UpdateUser(db, loggedInUser)
			} else {
				fmt.Println("Please login first!")
			}

		case 5: // hapus akun
			// mysql : DELETE FROM User WHERE username = ...
			fmt.Println("Masukkan Username yang ingin dihapus :")
			deleteAccount := entities.User{}
			fmt.Scanln(&deleteAccount.Username)

			deleteAccountRows, errDelete := db.Exec("DELETE FROM users WHERE username = ?", deleteAccount.Username)
			if err != nil {
				log.Fatal("error delete", errDelete.Error())
			} else {
				row, _ := deleteAccountRows.RowsAffected()
				if row > 0 {
					fmt.Println("Delete Success")
				} else {
					fmt.Println("Delete Failed")
				}
			}

		case 6: // fitur topup saldo
			fmt.Println("Top Up Balance")
			// loginUser := entities.User{}
			// fmt.Print("Masukkan Nomor Telepon : ")
			// fmt.Scanln(&loginUser.Phone)
			if loggedInUser.Id != "" {
				fmt.Print("Masukkan Password : ")
				fmt.Scanln(&loggedInUser.Password)

				// users, loggedin := controllers.Login(db, loggedInUser.Phone, loggedInUser.Password)
				// for _, v := range users {
				fmt.Println("Balance :", loggedInUser.Balance)
				controllers.AddMoney(db, loggedInUser)
			} else {
				fmt.Println("Please login first!")

			}

		case 7: // fitur transfer dana
			newTransfer := entities.Transfer{}
			fmt.Println("Masukan ID Pengirim :")
			fmt.Scanln(&newTransfer.User_id)
			fmt.Println("Masukan Receiver User ID:")
			fmt.Scanln(&newTransfer.Receiver_userid)
			fmt.Println("Masukan Nominal :")
			fmt.Scanln(&newTransfer.Amount)

			// Insert transfer data
			tranferResult, errTransfer := db.Exec("INSERT INTO TRANSFERS (user_id, receiver_userid, amount, status) VALUES (?, ?, ?, ?)", newTransfer.User_id, newTransfer.Receiver_userid, newTransfer.Amount, "SUCCESS")
			if errTransfer != nil {
				log.Fatal("error transfer", errTransfer.Error())
			} else {
				row, _ := tranferResult.RowsAffected()
				if row > 0 {
					fmt.Println("transfers success")
				} else {
					fmt.Println("transfer failed")
				}

				// Update receiver's balance
				updateReceiverBalance, errUpdateBalance := db.Exec("UPDATE users SET balance = balance + ? WHERE id = ?", newTransfer.Amount, newTransfer.Receiver_userid)
				if errUpdateBalance != nil {
					log.Fatal("error updating balance", errUpdateBalance.Error())
				} else {
					row, _ := updateReceiverBalance.RowsAffected()
					if row > 0 {
						fmt.Println("Update balance success")
					} else {
						fmt.Println("Update balance failed")
					}
				}

				// Update sender's balance
				updateSenderBalance, errUpdateSenderBalance := db.Exec("UPDATE users SET balance = balance - ? WHERE id = ?", newTransfer.Amount, newTransfer.User_id)
				if errUpdateSenderBalance != nil {
					log.Fatal("error updating sender's balance", errUpdateSenderBalance.Error())
				} else {
					row, _ := updateSenderBalance.RowsAffected()
					if row > 0 {
						fmt.Println("Update sender's balance success")
					} else {
						fmt.Println("Update sender's balance failed")
					}
				}
			}

		case 8: // fitur melihat history topup
			fmt.Println("Top-Up History")
			// loginUser := entities.User{}
			// fmt.Print("Masukkan Nomor Telepon : ")
			// fmt.Scanln(&loginUser.Phone)
			if loggedInUser.Id != "" {
				fmt.Print("Masukkan Password : ")
				fmt.Scanln(&loggedInUser.Password)
				// fmt.Println("==================")

				// users, loggedin := controllers.Login(db, loginUser.Phone, loginUser.Password)
				// for _, v := range users {
				// if loggedin {

				result := controllers.HistoryTopup(db, loggedInUser)
				for _, v := range result {
					fmt.Println("ID :", v.Topup_Id)
					fmt.Println("User_ID :", v.Id)
					fmt.Println("Name :", v.Name)
					fmt.Println("Amount :", v.Amount)
					fmt.Println("Status :", v.Status)
					fmt.Println("Updated Balance :", v.Balance+v.Amount)
					fmt.Println("Time :", v.Transaction_time_topup)
					fmt.Println("==================")
				}
			} else {
				fmt.Println("Please login first!")
			}

		case 9: // fitur melihat history transfer

			rows, errTranfserHistory := db.Query("SELECT user_id, receiver_userid, amount, status FROM transfers ORDER BY transaction_time")
			if errTranfserHistory != nil {
				log.Fatal("view transfer history failed", errTranfserHistory.Error())
			}
			var transferHistory []entities.Transfer
			for rows.Next() {
				var dataTransferHistory entities.Transfer
				errScan := rows.Scan(&dataTransferHistory.User_id, &dataTransferHistory.Receiver_userid, &dataTransferHistory.Amount, &dataTransferHistory.Status)
				if errScan != nil {
					log.Fatal("error scan", errScan.Error())
				}
				transferHistory = append(transferHistory, dataTransferHistory)
			}
			for _, v := range transferHistory {
				fmt.Println("Pengirim :", v.User_id)
				fmt.Println("Penerima :", v.Receiver_userid)
				fmt.Println("Jumlah :", v.Amount)
				fmt.Println("Status :", v.Status)
			}

		case 10: // fitur melihat profil user lain dengan menggunakan phone number
			viewByPhone := entities.User{}
			fmt.Println("Masukan Nomor Telepone :")
			fmt.Scanln(&viewByPhone.Phone)

			viewProfileRows, errViewProfile := db.Query("SELECT ID, username, name, phone FROM users WHERE phone = ?", viewByPhone.Phone)
			if err != nil {
				log.Fatal("error view profile", errViewProfile.Error())
			}
			var allViewUsers []entities.User
			for viewProfileRows.Next() {
				var dataViewUsers entities.User                                                                                        // variabel penampung untuk membaca viewProfileRows
				errScan := viewProfileRows.Scan(&dataViewUsers.Id, &dataViewUsers.Username, &dataViewUsers.Name, &dataViewUsers.Phone) // membaca yang dibaca oleh dbquery di viewProfileRows
				if err != nil {
					log.Fatal("error scan select", errScan.Error())
				}

				allViewUsers = append(allViewUsers, dataViewUsers)

			}
			fmt.Println("Data User :", allViewUsers)
		case 0: // fitur logout
			fmt.Println("Logging Out...")
			return

		default:
			{
				fmt.Println("Menu not available")
			}

		} // EOF menu
	}
}
