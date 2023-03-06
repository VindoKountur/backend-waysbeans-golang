package repositories

import (
	"backendwaysbeans/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	// GetTransaction(ID int) ([]models.TransactionRawQuery, error)
	GetTransaction(ID int) (models.Transaction, error)
	FindTransactionByUserID(ID int) ([]models.TransactionRawQuery, error)
	FindTransaction() ([]models.TransactionRawQuery, error)
	// FindTransaction() ([]models.Transaction, error)
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

func (h *repository) FindTransaction() ([]models.TransactionRawQuery, error) {
	// func (h *repository) FindTransaction() ([]models.Transaction, error) {
	var transactions []models.TransactionRawQuery
	// var transactions []models.Transaction
	err := h.db.Raw("SELECT t.id, t.user_id, u.name as fullname, u.email, t.status, t.created_at, t.updated_at, c.product_id, c.order_quantity, p.name as product_name, p.price, p.description, p.photo FROM transactions t INNER JOIN carts c ON c.transaction_id = t.id INNER JOIN products p ON p.id = c.product_id INNER JOIN users u ON u.id = t.user_id ORDER BY c.transaction_id ASC").Scan(&transactions).Error
	// err := h.db.Preload("User").Preload("cart").Find(&transactions).Error // Masih Error di Preload
	return transactions, err
}

func (h *repository) FindTransactionByUserID(UserID int) ([]models.TransactionRawQuery, error) {
	// func (h *repository) FindTransaction() ([]models.Transaction, error) {
	var transactions []models.TransactionRawQuery
	// var transactions []models.Transaction
	err := h.db.Raw("SELECT t.id, t.user_id, u.name as fullname, u.email, t.status, t.created_at, t.updated_at, c.product_id, c.order_quantity, p.name as product_name, p.price, p.description, p.photo FROM transactions t INNER JOIN carts c ON c.transaction_id = t.id INNER JOIN products p ON p.id = c.product_id INNER JOIN users u ON u.id = t.user_id WHERE t.user_id = ? ORDER BY c.transaction_id ASC", UserID).Scan(&transactions).Error
	// err := h.db.Preload("User").Preload("cart").Find(&transactions).Error // Masih Error di Preload
	return transactions, err
}

// func (h *repository) MinusProductTransaction(ID int) error {

// }
