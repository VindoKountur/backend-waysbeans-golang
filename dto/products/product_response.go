package productsdto

type ProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Stock       int    `json:"stock" form:"stock"`
	Photo       string `json:"photo" form:"photo"`
}
