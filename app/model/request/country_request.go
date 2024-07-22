package request

type StateRequestBaseOnCountry struct {
	CountryId int    `form:"countryId"`
	Page      int    `form:"page"`
	Search    string `form:"search"`
}
