package countrydto

type CreateCountryrRequest struct {
	NameCountry string `json:"country_name" form:"country_name" validate:"required"`
}

type UpdateCountryRequest struct {
	NameCountry string `json:"country_name" form:"country_name"`
}
