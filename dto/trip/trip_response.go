package tripdto

type TripResponse struct {
	ID           int    `json:"id"`
	Country_id   int    `json:"country_id" form:"country_id" validate:"required"`
	Title        string `json:"title" form:"title" validate:"required"`
	Accomodation string `json:"accomodation" form:"accomodation" validate:"required"`
	Transport    string `json:"transport" form:"transport" validate:"required"`
	Eat          string `json:"eat" form:"eat" validate:"required"`
	Day          int    `json:"day" form:"day" validate:"required"`
	Night        int    `json:"night" form:"night" validate:"required"`
	Date         string `json:"date" form:"date" validate:"required"`
	Price        int    `json:"price" form:"price" validate:"required"`
	Kuota        int    `json:"kuota" form:"kuota" validate:"required"`
	Description  string `json:"description" form:"description" validate:"required"`
	Image        string `json:"image" form:"image" validate:"required"`
}
type TripResponsedel struct {
	ID int `json:"id"`
}
