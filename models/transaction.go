package models

import "time"

type Transaction struct {
	ID         int          `json:"id" gorm:"primary_key:auto_increment"`
	Name       string       `json:"name" gorm:"type:varchar(255)"`
	Email      string       `json:"email" gorm:"type:varchar(255)"`
	Phone      string       `json:"phone" gorm:"type:varchar(255)"`
	Address    string       `json:"address" gorm:"type:varchar(255)"`
	PostCode   string       `json:"post_code" gorm:"type:varchar(255)"`
	UserID     int          `json:"user_id" gorm:"type:int"`
	User       UserResponse `json:"user"`
	TotalPrice int          `json:"total_price"`
	Status     string       `json:"status" gorm:"type:varchar(255)"`
	Cart       []Cart       `json:"products"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

type TransactionResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"date"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}

type TransactionRawQuery struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	Fullname      string    `json:"fullname"`
	Email         string    `json:"email"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	ProductID     int       `json:"product_id"`
	OrderQuantity int       `json:"order_quantity"`
	ProductName   string    `json:"product_name"`
	Price         int       `json:"price"`
	Description   string    `json:"description"`
	Photo         string    `json:"photo"`
}
