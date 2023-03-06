package handlers

import (
	dto "backendwaysbeans/dto/result"
	transactiondto "backendwaysbeans/dto/transaction"
	usersdto "backendwaysbeans/dto/users"
	"backendwaysbeans/models"
	"backendwaysbeans/repositories"
	"net/http"
	"strconv"

	// "github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type handlerUser struct {
	UserRepository        repositories.UserRepository
	TransactionRepository repositories.TransactionRepository
}

func HandlerUser(UserRepository repositories.UserRepository, TransactionRepository repositories.TransactionRepository) *handlerUser {
	return &handlerUser{UserRepository, TransactionRepository}
}

func (h *handlerUser) FindUsers(c echo.Context) error {
	users, err := h.UserRepository.FindUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: users})
}

func (h *handlerUser) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	userTransactionsRaw, err := h.TransactionRepository.FindTransactionByUserID(id)

	userTransaction := getTransactionListFromRaw(userTransactionsRaw)
	// user.Transaction = userTransaction

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertResponseUser(user, userTransaction)})
	// return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertResponseUser(user)})
}

func convertResponseUser(u models.User, transactions []transactiondto.TransactionResponse) usersdto.UserTransactionByIDResponse {
	return usersdto.UserTransactionByIDResponse{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		Profile:     u.Profile,
		Product:     u.Product,
		Addresses:   u.Addresses,
		Transaction: transactions,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}
