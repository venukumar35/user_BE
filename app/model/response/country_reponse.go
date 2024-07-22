package response

type CountryResponse struct {
	Id         uint   `json:"id" gorm:"column:id"`
	Name       string `json:"name" gorm:"column:name"`
	DialCode   string `json:"dialCode" gorm:"column:dialCode"`
	TotalCount int16  `json:"totalCount" gorm:"column:totalCount"`
}

type StateResponse struct {
	Id         uint   `json:"id" gorm:"column:id"`
	Name       string `json:"name" gorm:"column:name"`
	TotalCount int16  `json:"totalCount" gorm:"column:totalCount"`
}
