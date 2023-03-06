package repositories

import (
	"backendwaysbeans/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(email string) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) Login(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email =?", email).First(&user).Error
	return user, err
}
