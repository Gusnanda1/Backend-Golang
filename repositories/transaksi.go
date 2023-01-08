package repositories

import (
	"Backend/models"

	"gorm.io/gorm"
)

type TransaksiRepository interface {
	FindTransaksi() ([]models.Transaction, error)
	GetTransaksi(ID int) (models.Transaction, error)
	AddTransaksi(transaksi models.Transaction) (models.Transaction, error)
	UpdateTransaksi(status string, ID int) (models.Transaction, error)
	DeleteTransaksi(transaksi models.Transaction) (models.Transaction, error)
	GetOneTransaction(ID string) (models.Transaction, error)
}

func RepositoryTransaksi(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransaksi() ([]models.Transaction, error) {
	var Transaksi []models.Transaction
	err := r.db.Preload("Trip").Preload("Trip.Country").Preload("User").Find(&Transaksi).Error

	return Transaksi, err
}

func (r *repository) GetTransaksi(ID int) (models.Transaction, error) {
	var Transaksi models.Transaction
	err := r.db.Preload("Trip").Preload("Trip.Country").Preload("User").First(&Transaksi, ID).Error

	return Transaksi, err
}

func (r *repository) AddTransaksi(transaksi models.Transaction) (models.Transaction, error) {

	err := r.db.Create(&transaksi).Error
	return transaksi, err
}

func (r *repository) UpdateTransaksi(status string, ID int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("Trip").First(&transaction, ID)

	// If is different & Status is "success" decrement product quantity
	if status != transaction.Status && status == "success" {
		var trip models.Trip
		r.db.First(&trip, transaction.Trip.ID)
		trip.Kuota = trip.Kuota - 1
		r.db.Save(&trip)
	}

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaksi(transaksi models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transaksi).Error

	return transaksi, err
}
func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Trip").Preload("User").First(&transaction, "id = ?", ID).Error

	return transaction, err
}
