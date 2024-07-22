package request

type RequestForGetProduct struct {
	Page   int    `form:"page"`
	Search string `form:"search"`
}

type RequestProductById struct {
	Id     int    `form:"id"`
	Search string `form:"search"`
	Page   int    `form:"page"`
}
