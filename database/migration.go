package database

import (
	"Backend/models"
	"Backend/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Trip{},
		&models.Country{},
		&models.Transaction{},
	)
	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}
	fmt.Println("Migration Success")
}
