package handlers

import (
	dto "backendwaysbeans/dto/result"
	"backendwaysbeans/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type cartHandler struct {
	CartRepository repositories.CartRepository
}

func HandlerCart(CartRepository repositories.CartRepository) *cartHandler {
	return &cartHandler{
		CartRepository: CartRepository,
	}
}

func (h *cartHandler) FindCarts(c echo.Context) error {
	carts, err := h.CartRepository.FindCarts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: carts})
}
