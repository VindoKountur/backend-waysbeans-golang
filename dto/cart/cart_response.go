package cartdto

type CartResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	Description   string `json:"description"`
	Photo         string `json:"photo"`
	OrderQuantity int    `json:"orderQuantity"`
}
