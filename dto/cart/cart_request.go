package cartdto

type CreateCartRequest struct {
	ProductID     int `json:"id"`
	OrderQuantity int `json:"orderQuantity"`
}
