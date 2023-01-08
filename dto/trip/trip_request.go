package tripdto

import "time"

type TripRequest struct {
	Title        string    `json:"title" form:"title" validate:"required"`
	Country_id   int       `json:"country_id" form:"country_id"`
	Accomodation string    `json:"accomodation" form:"accomodation" `
	Transport    string    `json:"transport" form:"transport" `
	Eat          string    `json:"eat" form:"eat" `
	Day          int       `json:"day" form:"day" `
	Night        int       `json:"night" form:"night" `
	Date         time.Time `json:"date" form:"date" `
	Price        int       `json:"price" form:"price" `
	Kuota        int       `json:"kuota" form:"kuota" `
	Description  string    `json:"description" form:"description" `
	Image        string    `json:"image" form:"image"`
}

type TripUpdateRequest struct {
	Title        string    `json:"title" form:"title"`
	Country_id   int       `json:"country_id" form:"country_id"`
	Accomodation string    `json:"accomodation" form:"accomodation"`
	Transport    string    `json:"transport" form:"transport"`
	Eat          string    `json:"eat" form:"eat"`
	Day          int       `json:"day" form:"day"`
	Night        int       `json:"night" form:"night"`
	Date         time.Time `json:"date" form:"date"`
	Price        int       `json:"price" form:"price"`
	Kuota        int       `json:"kuota" form:"kuota"`
	Description  string    `json:"description" form:"description"`
	Image        string    `json:"image" form:"image"`
}
