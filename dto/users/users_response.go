package usersdto

import (
	transactiondto "backendwaysbeans/dto/transaction"
	"backendwaysbeans/models"
	"time"
)

type UserResponse struct {
	ID      int                    `json:"id"`
	Name    string                 `json:"name" form:"name" validate:"required"`
	Email   string                 `json:"email" form:"email" validate:"required"`
	Profile models.ProfileResponse `json:"profile"`
}

type UserTransactionResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UserTransactionByIDResponse struct {
	ID          int                                  `json:"id" gorm:"primary_key:auto_increment"`
	Email       string                               `json:"email" gorm:"type: varchar(255)"`
	Name        string                               `json:"name" gorm:"type: varchar(255)"`
	Profile     models.ProfileResponse               `json:"profile"`
	Product     []models.ProductUserResponse         `json:"products"`
	Addresses   []models.AddressesResponse           `json:"addresses"`
	Transaction []transactiondto.TransactionResponse `json:"transactions"`
	CreatedAt   time.Time                            `json:"-"`
	UpdatedAt   time.Time                            `json:"-"`
}
