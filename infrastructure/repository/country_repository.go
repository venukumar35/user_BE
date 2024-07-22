package repository

import (
	"TheBoys/app/model/request"
	"TheBoys/app/model/response"
	"TheBoys/domain"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func NewCountryRepository(db *gorm.DB) domain.CountryRepository {
	return &countryRepository{db}
}

type countryRepository struct {
	db *gorm.DB
}

func (r *countryRepository) FindCountry(req request.CommonRequest) ([]response.CountryResponse, error) {
	var limit int = 10

	var data *[]response.CountryResponse

	search := fmt.Sprintf("%%%s%%", strings.TrimSpace(req.Search))

	var baseQuery strings.Builder

	baseQuery.WriteString(`SELECT "id" as "id","name" AS "name","dialCode" AS "dialCode" ,COUNT(*) OVER (PARTITION BY 1) AS totalCount FROM "Country"`)

	if len(search) > 2 {
		baseQuery.WriteString(fmt.Sprintf(` WHERE "name" LIKE '%s'`, search))
	}

	if req.Page > 0 {
		baseQuery.WriteString(fmt.Sprintf(` LIMIT %d OFFSET %d`, limit, (req.Page)*limit))
	}

	err := r.db.Raw(baseQuery.String()).Scan(&data).Error

	if err != nil {
		return nil, err
	}

	return *data, nil
}

func (r *countryRepository) FindState(req request.StateRequestBaseOnCountry) ([]response.StateResponse, error) {
	var limit int = 10

	var data *[]response.StateResponse

	search := fmt.Sprintf("%%%s%%", strings.TrimSpace(req.Search))

	var baseQuery strings.Builder

	baseQuery.WriteString(fmt.Sprintf(`SELECT "id" as "id","name" AS "name",COUNT(*) OVER (PARTITION BY 1) AS "totalCount" FROM "State" WHERE "countryId" = %d `, req.CountryId))

	if len(search) > 2 {
		baseQuery.WriteString(fmt.Sprintf(` WHERE "name" LIKE '%s'`, search))
	}

	if req.Page > 0 {
		baseQuery.WriteString(fmt.Sprintf(` LIMIT %d OFFSET %d`, limit, (req.Page)*limit))
	}

	err := r.db.Raw(baseQuery.String()).Scan(&data).Error

	if err != nil {
		return nil, err
	}

	return *data, nil
}
