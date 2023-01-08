package models

import "time"

type Trip struct {
	ID           int                 `json:"id"  gorm:"primary_key:auto_increment"`
	Title        string              `json:"title" gorm:"type : varchar (255)"`
	Country_id   int                 `json:"-"`
	Country      CountrytripResponse `json:"country" `
	Accomodation string              `json:"accomodation" gorm:"type : varchar (255)"`
	Transport    string              `json:"transport" gorm:"type : varchar (255)"`
	Eat          string              `json:"eat" gorm:"type : varchar (255)"`
	Day          int                 `json:"day" gorm:"type : varchar (255)"`
	Night        int                 `json:"night" gorm:"type : varchar (255)"`
	Date         time.Time           `json:"date" `
	Price        int                 `json:"price" gorm:"type : int"`
	Kuota        int                 `json:"kuota" gorm:"type : int"`
	Description  string              `json:"description" gorm:"type : varchar (255)"`
	Image        string              `json:"image" gorm:"type : varchar (255)"`
}

type TripResponse struct {
	ID           int                 `json:"id"`
	Title        string              `json:"title"`
	Country_id   int                 `json:"country_id"`
	Country      CountrytripResponse `json:"country"`
	Accomodation string              `json:"accomodation"`
	Transport    string              `json:"transport"`
	Eat          string              `json:"eat"`
	Day          int                 `json:"day"`
	Night        int                 `json:"night"`
	Date         time.Time           `json:"date"`
	Price        int                 `json:"price"`
	Kuota        int                 `json:"kuota"`
	Description  string              `json:"description"`
	Image        string              `json:"image"`
}

func (TripResponse) TableName() string {
	return "trips"
}
