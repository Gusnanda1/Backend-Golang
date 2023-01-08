package repositories

import (
	"Backend/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUser() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	GetUserEmail(email string) (models.User, error)
	AddUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUser() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var users models.User
	err := r.db.First(&users, ID).Error

	return users, err
}

func (r *repository) GetUserEmail(email string) (models.User, error) {
	var users models.User
	err := r.db.First(&users, email).Error

	return users, err
}

func (r *repository) AddUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
