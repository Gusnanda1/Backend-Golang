package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaInit() {
	var err error
	dsn := "http://root:3D4ZhsUSgSORbMIYlonY@tcp(containers-us-west-85.railway.app:8065)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected To Database")
}
