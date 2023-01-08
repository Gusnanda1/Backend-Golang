package repositories

import (
	"Backend/models"

	"gorm.io/gorm"
)

type TripRepository interface {
	FindTrip() ([]models.Trip, error)
	GetTrip(ID int) (models.Trip, error)
	Createtrip(trip models.Trip) (models.Trip, error)
	UpdatedTrip(trip models.Trip) (models.Trip, error)
	DeleteTrip(trip models.Trip) (models.Trip, error)
}

func RepositoryTrip(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTrip() ([]models.Trip, error) {
	var Trip []models.Trip
	err := r.db.Preload("Country").Find(&Trip).Error

	return Trip, err
}

func (r *repository) GetTrip(ID int) (models.Trip, error) {
	var Trip models.Trip
	err := r.db.Preload("Country").First(&Trip, ID).Error

	return Trip, err
}

func (r *repository) Createtrip(trip models.Trip) (models.Trip, error) {
	err := r.db.Create(&trip).Error

	return trip, err
}

func (r *repository) UpdatedTrip(trip models.Trip) (models.Trip, error) {
	// err := r.db.Exec("UPDATE trips SET title=?, country_id=?, accomodation=?, transport=?, eat=?, day=?, night=?, date=?, price=?, kuota=?, description=?, image=? WHERE trips.id=?", trip.Title, trip.Country_id, trip.Accomodation, trip.Transport, trip.Eat, trip.Day, trip.Night, trip.Date, trip.Price, trip.Kuota, trip.Description, trip.Image, trip.ID).Error

	err := r.db.Model(&trip).Updates(trip).Error

	return trip, err
}

func (r *repository) DeleteTrip(trip models.Trip) (models.Trip, error) {
	err := r.db.Delete(&trip).Error

	return trip, err
}
