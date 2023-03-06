package transactiondto

import (
	cartdto "backendwaysbeans/dto/cart"
	"time"
)

type TransactionResponse struct {
	ID        int                     `json:"id"`
	UserID    int                     `json:"-"`
	User      UserTransactionResponse `json:"user"`
	Status    string                  `json:"status"`
	Cart      []cartdto.CartResponse  `json:"products"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
}

type UserTransactionResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}
