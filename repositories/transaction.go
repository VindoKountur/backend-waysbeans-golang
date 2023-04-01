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
	UpdateTransaction(status string, orderId int) (models.Transaction, error)
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
	err := h.db.Where("id = ?", ID).Preload("Cart.Product").Preload("User").First(&transaction).Error
	return transaction, err
}

func (h *repository) FindTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := h.db.Preload("User").Preload("Cart.Product").Order("created_at desc").Find(&transactions).Error
	return transactions, err
}

func (h *repository) FindTransactionByUserID(UserID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := h.db.Preload("User").Preload("Cart.Product").Where("user_id = ?", UserID).Find(&transactions).Error
	return transactions, err
}

func (r *repository) UpdateTransaction(status string, orderId int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("Cart").First(&transaction, orderId)

	if status != transaction.Status && status == "success" {
		for _, cart := range transaction.Cart {
			var product models.Product
			r.db.First(&product, cart.ProductID)
			product.Stock -= cart.OrderQuantity
			if product.Stock < 0 {
				product.Stock = 0
			}
			r.db.Save(&product)
		}
	}

	transaction.Status = status
	err := r.db.Save(&transaction).Error
	return transaction, err
}
