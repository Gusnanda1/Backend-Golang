package models

type User struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	Fullname string `json:"fullname" gorm:"type : varchar (255)"`
	Email    string `json:"email" gorm:"type : varchar (255)"`
	Password string `json:"password" gorm:"type : varchar (255)"`
	Gender   string `json:"gender" gorm:"type:varchar(255)"`
	Phone    string `json:"phone" gorm:"type : varchar (255)"`
	Address  string `json:"address" gorm:"type : varchar (255)"`
	Role     string `json:"role" gorm:"type : varchar (255)"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Role     string `json:"role"`
}

func (UserResponse) TableName() string {
	return "users"
}
