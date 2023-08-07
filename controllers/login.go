package controllers

import (
	"be18/acc-service/entities"
	"database/sql"
	"log"
)

func Login(db *sql.DB, phone, password string) ([]entities.User, bool) {

	rows, errLogin := db.Query("SELECT id,name,username,address,email,phone,password,balance FROM users WHERE phone=? AND password=?", phone, password)
	if errLogin != nil {
		log.Fatal("Error query select", errLogin.Error())
	}
	// variabel untuk menyimpan semua data yang dibaca di db.Query
	var allUsers []entities.User
	// membaca per baris
	for rows.Next() {
		var dataUserRow entities.User // variabel untuk menyimpan data per baris
		errScan := rows.Scan(&dataUserRow.Id, &dataUserRow.Name, &dataUserRow.Username, &dataUserRow.Address, &dataUserRow.Email, &dataUserRow.Phone, &dataUserRow.Password, &dataUserRow.Balance)
		if errScan != nil {
			log.Fatal("error scan select", errScan.Error())
		}
		// menambahkan tiap baris data yang dibaca ke slice
		allUsers = append(allUsers, dataUserRow)
	}

	return allUsers, len(allUsers) > 0 //login berhasil
}
