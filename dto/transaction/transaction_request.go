package transactiondto

import cartdto "backendwaysbeans/dto/cart"

type CreateTransactionRequest struct {
	Name     string                      `json:"name" validate:"required"`
	Email    string                      `json:"email" validate:"required"`
	Phone    string                      `json:"phone" validate:"required"`
	Address  string                      `json:"address" validate:"required"`
	PostCode string                      `json:"postCode" validate:"required"`
	Products []cartdto.CreateCartRequest `json:"products" validate:"required"`
}
