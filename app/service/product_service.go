package service

import (
	"TheBoys/app/model/request"
	"TheBoys/domain"
	"TheBoys/utills"
)

func NewProductService(repo domain.ProductRepository) domain.ProductService {
	return &ProductServices{repo}
}

type ProductServices struct {
	repo domain.ProductService
}

func (s *ProductServices) GetProducts() (*utills.PaginationResponse, error) {
	data, err := s.repo.GetProducts()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *ProductServices) GetProductCategory() (interface{}, error) {
	data, err := s.repo.GetProductCategory()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *ProductServices) GetProductById(req request.RequestProductById) (*utills.PaginationResponse, error) {
	data, err := s.repo.GetProductById(req)

	if err != nil {
		return nil, err
	}

	return data, nil
}
