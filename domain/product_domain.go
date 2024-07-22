package domain

import (
	"TheBoys/app/model/request"
	"TheBoys/utills"
)

type ProductService interface {
	GetProducts() (*utills.PaginationResponse, error)
	GetProductById(req request.RequestProductById) (*utills.PaginationResponse, error)
	GetProductCategory() (interface{}, error)
}
type ProductRepository interface {
	GetProducts() (*utills.PaginationResponse, error)
	GetProductById(req request.RequestProductById) (*utills.PaginationResponse, error)
	GetProductCategory() (interface{}, error)
}
