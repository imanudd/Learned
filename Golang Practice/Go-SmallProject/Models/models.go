package models

type Customers struct {
	Id           int
	Name         string
	Phone_number string
}

type Cassiers struct {
	Id   int
	Name string
}

type Services struct {
	Id          int
	ServiceName string
	Unit        string
	Price       int
}

type Transaction struct {
	Id         int
	IdCustomer int
	IdCassier  int
	StartDate  string
	EndDate    string
}

type Detail_Transaction struct {
	Id            int
	Idtransaction int
	IdService     int
	Qty           int
}
