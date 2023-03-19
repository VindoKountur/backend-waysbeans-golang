package repositories

import (
	"backendwaysbeans/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProduct() ([]models.Product, error)
	FindProductByName(name string) ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(product models.Product) (models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProduct() ([]models.Product, error) {
	var products []models.Product
	// err := r.db.Raw("SELECT * FROM products").Scan(&products).Error
	err := r.db.Preload("User").Find(&products).Error
	return products, err
}
func (r *repository) FindProductByName(name string) ([]models.Product, error) {
	var products []models.Product
	searchName := "%" + name + "%"
	// err := r.db.Raw("SELECT * FROM products").Scan(&products).Error
	err := r.db.Where("name LIKE BINARY ?", searchName).Preload("User").Find(&products).Error
	return products, err
}

func (r *repository) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	// err := r.db.Raw("SELECT * FROM products WHERE id = ?", ID).Scan(&product).Error
	err := r.db.Preload("User").First(&product, ID).Error
	return product, err
}

func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	// err := r.db.Exec("INSERT INTO product (name, price, description, stock, photo, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?)", product.Name, product.Price, product.Description, product.Stock, product.Photo, time.Now(), time.Now()).Error
	err := r.db.Create(&product).Error

	return product, err
}

func (r *repository) UpdateProduct(product models.Product) (models.Product, error) {
	// err := r.db.Raw("UPDATE product SET name = ?, price = ?, description = ?, stock = ?, photo = ?, updated_at = ? WHERE id = ?", product.Name, product.Price, product.Description, product, product.Stock, product.Photo, time.Now(), ID).Error
	err := r.db.Save(&product).Error
	return product, err
}

func (r *repository) DeleteProduct(product models.Product) (models.Product, error) {
	// err := r.db.Exec("DELETE FROM product WHERE id = ?", ID).Error
	err := r.db.Delete(&product).Error
	return product, err
}
