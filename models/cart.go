package models

type Cart struct {
	ID            int                 `json:"id" gorm:"primary_key:auto_increment"`
	OrderQuantity int                 `json:"order_quantity" gorm:"type:int"`
	TransactionID int                 `json:"transaction_id" gorm:"type:int"`
	Transaction   TransactionResponse `json:"-"`
	ProductID     int                 `json:"product_id"`
	Product       ProductCartResponse `json:"product"`
}

type CartResponse struct {
	OrderQuantity int                 `json:"order_quantity"`
	ProductID     int                 `json:"product_id"`
	Product       ProductCartResponse `json:"product"`
	TransactionID int                 `json:"-"`
}

func (CartResponse) TableName() string {
	return "carts"
}
