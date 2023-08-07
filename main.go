package main

import (
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
	case 4: // update profil
	case 5: // hapus akun
	case 6: // fitur topup saldo
	case 7: // fitur transfer dana
	case 8: // fitur melihat history topup
	case 9: // fitur melihat history transfer
	case 10: // fitur melihat profil user lain dengan menggunakan phone number
	case 0: // fitur logout

	} // EOF menu

}
