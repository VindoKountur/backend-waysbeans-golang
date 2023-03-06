package addressdto

type CreateAddressRequest struct {
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Address  string `json:"address" validate:"required"`
	PostCode string `json:"post_code" validate:"required"`
}

type UpdateAddressRequest struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	PostCode string `json:"post_code"`
}
