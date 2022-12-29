package product

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Product struct {
	Id          int    `json:"id"`
	Category    string `json:"category"`
	ProductName string `json:"product_name"`
	Weight      string `json:"weight"`
	Price       int    `json:"price"`
}
