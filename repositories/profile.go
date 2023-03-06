package repositories

import (
	"backendwaysbeans/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetProfileByUser(userID int) (models.Profile, error)
	CreateProfileByUser(profile models.Profile) (models.Profile, error)
	UpdateProfileByUser(profile models.Profile, UserID int) (models.Profile, error)
}

func RepositoryProfile(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (h *repository) GetProfileByUser(userID int) (models.Profile, error) {
	var profile models.Profile
	err := h.db.Where("user_id =?", userID).First(&profile).Error
	return profile, err
}

func (h *repository) CreateProfileByUser(profile models.Profile) (models.Profile, error) {
	err := h.db.Create(&profile).Error
	return profile, err
}

func (h *repository) UpdateProfileByUser(profile models.Profile, UserID int) (models.Profile, error) {
	err := h.db.Where("user_id =?", UserID).Save(&profile).Error
	return profile, err
}
