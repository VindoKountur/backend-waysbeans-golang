package repositories

import (
	"backendwaysbeans/models"

	"gorm.io/gorm"
)

type AddressRepository interface {
	FindAddress() ([]models.Addresses, error)
	FindAddressByUser(User_ID int) ([]models.Addresses, error)
	GetAddress(ID int) (models.Addresses, error)
	CreateAddress(models.Addresses) (models.Addresses, error)
	UpdateAddress(models.Addresses) (models.Addresses, error)
	DeleteAddress(models.Addresses, int) (models.Addresses, error)
}

func RepositoryAddress(db *gorm.DB) *repository {
	return &repository{db}
}

func (h *repository) FindAddress() ([]models.Addresses, error) {
	var address []models.Addresses
	err := h.db.Find(&address).Error
	return address, err
}

func (h *repository) FindAddressByUser(User_ID int) ([]models.Addresses, error) {
	var addresses []models.Addresses
	err := h.db.Where("user_id = ?", User_ID).Find(&addresses).Error
	return addresses, err
}

func (h *repository) GetAddress(ID int) (models.Addresses, error) {
	var address models.Addresses
	err := h.db.First(&address, ID).Error
	return address, err
}

func (h *repository) CreateAddress(address models.Addresses) (models.Addresses, error) {
	err := h.db.Preload("User").Create(&address).Error
	return address, err
}

func (h *repository) UpdateAddress(address models.Addresses) (models.Addresses, error) {
	err := h.db.Save(&address).Error
	return address, err
}

func (h *repository) DeleteAddress(address models.Addresses, AddressID int) (models.Addresses, error) {
	err := h.db.Delete(&address, AddressID).Error
	return address, err
}
