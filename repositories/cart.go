package repositories

import (
	"backendwaysbeans/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	CreateCart(cart models.Cart) (models.Cart, error)
	DeleteCart(cart models.Cart) (models.Cart, error)
	FindCartByTransactionID(transactionID int) ([]models.Cart, error)
	// DEBUG
	FindCarts() ([]models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (h *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := h.db.Create(&cart).Error
	return cart, err
}

func (h *repository) DeleteCart(cart models.Cart) (models.Cart, error) {
	err := h.db.Delete(&cart).Error
	return cart, err
}

func (h *repository) FindCartByTransactionID(transactionID int) ([]models.Cart, error) {
	var carts []models.Cart
	err := h.db.Where("transaction_id =?", transactionID).Preload("Product").Find(&carts).Error
	return carts, err
}

// DEBUG
func (h *repository) FindCarts() ([]models.Cart, error) {
	var carts []models.Cart
	err := h.db.Preload("Product").Find(&carts).Error
	return carts, err
}
