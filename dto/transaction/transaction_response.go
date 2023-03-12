package transactiondto

import (
	cartdto "backendwaysbeans/dto/cart"
	"time"
)

type TransactionResponse struct {
	ID        int                     `json:"id"`
	UserID    int                     `json:"-"`
	User      UserTransactionResponse `json:"user"`
	Name      string                  `json:"name"`
	Email     string                  `json:"email"`
	Phone     string                  `json:"phone"`
	Address   string                  `json:"address"`
	PostCode  string                  `json:"post_code"`
	Status    string                  `json:"status"`
	CreatedAt time.Time               `json:"created_at"`
	Cart      []cartdto.CartResponse  `json:"products"`
}

type UserTransactionResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}
