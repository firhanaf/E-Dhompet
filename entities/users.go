package entities

type User struct {
	Id                     string
	Username               string
	Name                   string
	Email                  string
	Password               string
	Address                string
	Phone                  string
	Balance                int64
	Amount                 int64
	Status                 string
	Topup_Id               int64
	Transaction_time_topup string
}

type Transfer strcut {
	Id					string
	User_Id 			string
	Receiver_userId 	string
	Amount 				int
	Status 				string
	Transaction_time 	string
}