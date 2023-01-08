package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaInit() {
	var err error
	dsn := "root:KhDhTPy66BSkv0Rr7g8r@tcp(containers-us-west-107.railway.app:5572/railway)?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected To Database")
}
