package service

import (
	"TheBoys/app/model/request"
	"TheBoys/domain"
	"TheBoys/utills"
)

func NewCountryService(repo domain.CountryRepository) domain.CountryService {
	return &CountryService{repo}
}

type CountryService struct {
	repo domain.CountryRepository
}

func (s *CountryService) FindCountry(req request.CommonRequest) (*utills.PaginationResponse, error) {
	data, err := s.repo.FindCountry(req)
	if err != nil {
		return nil, err
	}

	if len(data) > 0 {

		count := data[0].TotalCount

		response := utills.PaginatedResponse(int64(count), req.Page, data)

		return response, nil
	}

	return nil, nil

}

// FindState implements domain.CountryService.
func (s *CountryService) FindState(req request.StateRequestBaseOnCountry) (*utills.PaginationResponse, error) {
	data, err := s.repo.FindState(req)
	if err != nil {
		return nil, err
	}

	if len(data) > 0 {

		count := data[0].TotalCount

		response := utills.PaginatedResponse(int64(count), req.Page, data)

		return response, nil
	}

	return nil, nil
}
