package countrydto

type CountryResponse struct {
	ID          int    `json:"id"`
	CountryName string `json:"name_country" form:"name_country" validate:"required"`
}

type CountryResponseDel struct {
	ID int `json:"id"`
}
