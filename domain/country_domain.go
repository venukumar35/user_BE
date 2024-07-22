package domain

import (
	"TheBoys/app/model/request"
	"TheBoys/app/model/response"
	"TheBoys/utills"
)

type CountryService interface {
	FindCountry(req request.CommonRequest) (*utills.PaginationResponse, error)
	FindState(req request.StateRequestBaseOnCountry) (*utills.PaginationResponse, error)
}
type CountryRepository interface {
	FindCountry(req request.CommonRequest) ([]response.CountryResponse, error)
	FindState(req request.StateRequestBaseOnCountry) ([]response.StateResponse, error)
}
