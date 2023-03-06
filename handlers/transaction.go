package handlers

import (
	dto "backendwaysbeans/dto/result"
	transactiondto "backendwaysbeans/dto/transaction"
	"backendwaysbeans/models"
	"backendwaysbeans/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
	CartRepository        repositories.CartRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository, CartRepository repositories.CartRepository) *handlerTransaction {
	return &handlerTransaction{
		TransactionRepository: TransactionRepository,
		CartRepository:        CartRepository,
	}
}

func (h *handlerTransaction) CreateTransaction(c echo.Context) error {
	request := new(transactiondto.CreateTransactionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	idUserLogin := (c.Get("userLogin").(jwt.MapClaims)["id"]).(float64)

	newTransactions := models.Transaction{
		UserID:    int(idUserLogin),
		Name:      request.Name,
		Email:     request.Email,
		Phone:     request.Phone,
		Status:    "Waiting approved",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdTransaction, err := h.TransactionRepository.CreateTransaction(newTransactions)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	for _, product := range request.Products {
		newCart := models.Cart{
			TransactionID: createdTransaction.ID,
			ProductID:     product.ProductID,
			OrderQuantity: product.OrderQuantity,
		}
		_, err := h.CartRepository.CreateCart(newCart)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		}
		// Kurangi Stock Product
	}

	data, err := h.TransactionRepository.GetTransaction(createdTransaction.ID)
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: (data)})
}

func (h *handlerTransaction) FindTransaction(c echo.Context) error {
	transactions, err := h.TransactionRepository.FindTransaction()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: transactions})
}

func (h *handlerTransaction) GetTransaction(c echo.Context) error {
	transactionID, _ := strconv.Atoi(c.Param("id"))
	transactionsRaw, err := h.TransactionRepository.GetTransaction(transactionID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: (transactionsRaw)})
	// return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertTransactionRawToResponse(transactionsRaw)})
}

func convertCartResponse(c []models.Cart) []models.CartResponse {
	var cartResponse []models.CartResponse
	for _, cart := range c {
		cartResponse = append(cartResponse, models.CartResponse{
			ProductID:     cart.ProductID,
			OrderQuantity: cart.OrderQuantity,
		})
	}
	return cartResponse
}
