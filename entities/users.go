package entities

type User struct {
	Id                         string
	Username                   string
	Name                       string
	Email                      string
	Password                   string
	Address                    string
	Phone                      string
	Balance                    int64
	Amount                     int64
	Status                     string
	Topup_Id                   int64
	Transaction_time_topup     string
	Transaction_time_transfers string
}

type Transfer struct {
	Id              string
	User_id         string
	Receiver_userid string
	Amount          int64
	Status          string
}
