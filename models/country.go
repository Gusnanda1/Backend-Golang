package models

type Country struct {
	ID          int    `json:"id"  gorm:"primary_key:auto_increment"`
	NameCountry string `json:"name_country" gorm:"type : varchar (255)"`
}

type CountrytripResponse struct {
	ID          int    `json:"id"`
	NameCountry string `json:"name_country"`
}

func (CountrytripResponse) TableName() string {
	return "countries"
}
