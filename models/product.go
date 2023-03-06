package models

import "time"

type Product struct {
	ID          int            `json:"id" gorm:"primary_key:auto_increment"`
	Name        string         `json:"name" gorm:"type: varchar(255)"`
	Price       int            `json:"price" gorm:"type: int"`
	Stock       int            `json:"stock" gorm:"type: int"`
	Description string         `json:"description" gorm:"type: text"`
	Photo       string         `json:"photo" gorm:"type: varchar(255)"`
	UserID      int            `json:"-"`
	User        UserResponse   `json:"user"`
	Cart        []CartResponse `json:"-"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type ProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Photo       string `json:"photo"`
}

func (ProductResponse) TableName() string {
	return "products"
}

type ProductUserResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Photo  string `json:"photo"`
	Price  int    `json:"price"`
	Stock  int    `json:"stock"`
	UserID int    `json:"-"`
}

func (ProductUserResponse) TableName() string {
	return "products"
}

type ProductCartResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

func (ProductCartResponse) TableName() string {
	return "products"
}
