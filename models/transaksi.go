package models

type Transaction struct {
	ID         int          `json:"id"  gorm:"primary_key:auto_increment"`
	CounterQTY int          `json:"counterqty" gorm:"type : int"`
	Total      int          `json:"total" gorm:"type : int"`
	Status     string       `json:"status" gorm:"type : varchar (255)"`
	Image      string       `json:"image" gorm:"type : varchar (255)"`
	Trip_id    int          `json:"trip_id"`
	Trip       TripResponse `json:"trip"`
	UserID     int          `json:"user_id"`
	User       UserResponse `json:"user"`
}
