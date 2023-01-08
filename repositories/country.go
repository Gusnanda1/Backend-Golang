package repositories

import (
	"Backend/models"

	"gorm.io/gorm"
)

type CountryRepository interface {
	FindCountry() ([]models.Country, error)
	GetCountry(ID int) (models.Country, error)
	AddCountry(country models.Country) (models.Country, error)
	UpdateCountry(country models.Country) (models.Country, error)
	DeleteCountry(country models.Country) (models.Country, error)
}

func RepositoryCountry(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCountry() ([]models.Country, error) {
	var Country []models.Country
	err := r.db.Find(&Country).Error

	return Country, err
}

func (r *repository) GetCountry(ID int) (models.Country, error) {
	var Country models.Country
	err := r.db.First(&Country, ID).Error

	return Country, err
}

func (r *repository) AddCountry(country models.Country) (models.Country, error) {
	err := r.db.Create(&country).Error
	return country, err
}

func (r *repository) UpdateCountry(country models.Country) (models.Country, error) {
	err := r.db.Save(&country).Error

	return country, err
}

func (r *repository) DeleteCountry(country models.Country) (models.Country, error) {
	err := r.db.Delete(&country).Error

	return country, err
}
