package main

import (
	"fmt"
	"log"

	"be18/acc-service/controllers"
	"be18/acc-service/entities"

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
			fmt.Println("Masukkan ID :") // Email
			fmt.Scanln(&registerUser.Id)
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

			registerRows, errRegister := db.Exec("INSERT INTO Users (id, name, username, email, phone, password, address, balance) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", registerUser.Id, registerUser.Name, registerUser.Username, registerUser.Email, registerUser.Phone, registerUser.Password, registerUser.Address, 0)
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

			users, login := controllers.Login(db, loginUser.Phone, loginUser.Password)
			for _, v := range users {
				if login {
					fmt.Println("Login Success!\nWelcome", v.Name)
				} else {
					fmt.Println("Invalid phone or password")
				}
			}

		case 3: // view profile yang telah login
			fmt.Println("View Profile")
			loginUser := entities.User{}
			fmt.Print("Masukkan Nomor Telepon : ")
			fmt.Scanln(&loginUser.Phone)
			fmt.Print("Masukkan Password : ")
			fmt.Scanln(&loginUser.Password)

			users, loggedin := controllers.Login(db, loginUser.Phone, loginUser.Password)
			for _, v := range users {
				if loggedin {
					fmt.Println("=========================")
					fmt.Println("Displaying Profile of:", v.Name)
					fmt.Println("=========================")
					controllers.ReadProfile(db, v)
				} else {
					fmt.Println("Invalid phone or password")
				}
			}

		case 4: // update profil
		case 5: // hapus akun
		case 6: // fitur topup saldo
			fmt.Println("Top Up Balance")
			loginUser := entities.User{}
			fmt.Print("Masukkan Nomor Telepon : ")
			fmt.Scanln(&loginUser.Phone)
			fmt.Print("Masukkan Password : ")
			fmt.Scanln(&loginUser.Password)

			users, loggedin := controllers.Login(db, loginUser.Phone, loginUser.Password)
			for _, v := range users {
				if loggedin {
					fmt.Println("Balance :", v.Balance)
					controllers.AddMoney(db, v)
				}
			}

		case 7: // fitur transfer dana
		case 8: // fitur melihat history topup
		case 9: // fitur melihat history transfer
		case 10: // fitur melihat profil user lain dengan menggunakan phone number
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
