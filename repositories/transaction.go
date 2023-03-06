package repositories

import (
	"backendwaysbeans/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	FindTransactionByUserID(ID int) ([]models.Transaction, error)
	FindTransaction() ([]models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (h *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := h.db.Create(&transaction).Error
	return transaction, err
}

func (h *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := h.db.Where("id = ?", ID).Preload("Cart.Product").First(&transaction).Error
	return transaction, err
}

func (h *repository) FindTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := h.db.Preload("User").Preload("cart").Find(&transactions).Error // Masih Error di Preload
	return transactions, err
}

func (h *repository) FindTransactionByUserID(UserID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := h.db.Preload("User").Preload("cart").Find(&transactions).Error
	return transactions, err
}

// func (h *repository) MinusProductTransaction(ID int) error {

// }
