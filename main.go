package main

import (
	"be18/acc-service/entities"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       string
	Username string
	Name     string
	Email    string
	Password string
	Address  string
	Phone    string
}

func main() {

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

	//Main Menu
	fmt.Println("Selamat Datang di E-Dhompet")
	fmt.Println("Pilih Menu: \n 1. Register \n 2. Login \n 3. View Profile \n 4. Update Profile \n 5. Delete Account \n 6. Top Up \n 7. Transfer \n 8. Top Up History \n 9. Transfer History \n 10. View Profile by Phone Number \n 0. Log Out")
	var menu int
	fmt.Println("input pilihan anda:")
	fmt.Scanln(&menu)

	switch menu {
	case 1: // fitur Register

		// biasanya urutannya email, username, password, nama lengkap, nomor telepon, address

		registerUser := User{}
		fmt.Println("Masukkan Email :") // Email
		fmt.Scanln(&registerUser.Email)
		fmt.Println("Masukkan Username :") // Username
		fmt.Scanln(&registerUser.Username)
		fmt.Println("Masukkan Password :") // Password
		fmt.Scanln(&registerUser.Password)
		fmt.Println("Masukkan Username :") // Nama Lengkap
		fmt.Scanln(&registerUser.Name)
		fmt.Println("Masukkan Username :") // Nomor Telepon
		fmt.Scanln(&registerUser.Phone)
		fmt.Println("Masukkan Adress :") // Address
		fmt.Scanln(&registerUser.Address)

		registerRows, errRegister := db.Exec("INSERT INTO Users (id, name, username, email, phone, password, address) VALUES (?, ?, ?, ?, ?, ?, ?)", registerUser.Id, registerUser.Name, registerUser.Username, registerUser.Email, registerUser.Phone, registerUser.Password, registerUser.Address)
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
		loginUser := User{}
		fmt.Println("Masukkan Nomor Telepon :")
		fmt.Scanln(&loginUser.Phone)
		fmt.Println("Masukkan Password :")
		fmt.Scanln(&loginUser.Password)

		rows, errLogin := db.Query("select phone, password from users where phone = ? AND password = ?")
		if errLogin != nil { // ketika terjadi error saat menjalankan select
			log.Fatal("error run query select", errLogin.Error())
		}

		var allUsers []User

	case 3: // view profile yang telah login
		// mysql : SELECT * FROM USER WHERE ID = ...
		viewProfileRows, errViewProfile := db.Query("SELECT id, username, name, phone FROM users")

		if err != nil {
			log.Fatal("error view profile", errViewProfile.Error())
		}
		var allViewUsers []User
		for viewProfileRows.Next() {
			var dataViewUsers User
			errScan := viewProfileRows.Scan(&dataViewUsers.Id, &dataViewUsers.Username, &dataViewUsers.Name, dataViewUsers.Phone)
			if err != nil {
				log.Fatal("error scan select", errScan.Error())
			}
			allViewUsers = append(allViewUsers, dataViewUsers)
		}

	case 4: // update profil
	case 5: // hapus akun
	// mysql : DELETE FROM User WHERE ID = ...

	case 6: // fitur topup saldo
	case 7: // fitur transfer dana
	case 8: // fitur melihat history topup
	case 9: // fitur melihat history transfer
		newTransfer := entities.Transfer{}
		fmt.Println("Masukan jumlah uang :")
		fmt.Scanln(&newTransfer.Amount)
		fmt.Println("Masukan nomor penerima :")
		fmt.Scanln(&)

	case 10: // fitur melihat profil user lain dengan menggunakan phone number
	case 0: // fitur logout

	} // EOF menu

}
