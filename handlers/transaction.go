package handlers

import (
	cartdto "backendwaysbeans/dto/cart"
	dto "backendwaysbeans/dto/result"
	transactiondto "backendwaysbeans/dto/transaction"
	"backendwaysbeans/models"
	"backendwaysbeans/repositories"
	"net/http"
	"sort"
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

	// dataRaw, err := h.TransactionRepository.GetTransaction(createdTransaction.ID)
	data, err := h.TransactionRepository.GetTransaction(createdTransaction.ID)
	// carts, err := h.CartRepository.FindCartByTransactionID(data.ID)
	// fmt.Println(carts)
	// return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertTransactionRawToResponse(dataRaw)})

	// return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertTransactionResponse(data, carts)})
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: (data)})
}

func (h *handlerTransaction) FindTransaction(c echo.Context) error {
	transactions, err := h.TransactionRepository.FindTransaction()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	transactionList := getTransactionListFromRaw(transactions)

	// for idx, transaction := range transactions {
	// 	cartByID, err := h.CartRepository.FindCartByTransactionID(transaction.ID)
	// 	if err != nil {
	// 		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	// 	}
	// 	transactions[idx].Cart = convertCartResponse(cartByID)
	// }

	// carts, err := h.CartRepository.FindCartByTransactionID(transactions.)
	// fmt.Println(carts)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: transactionList})
	// return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: transactions})
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

func getTransactionListFromRaw(transactions []models.TransactionRawQuery) []transactiondto.TransactionResponse {
	sort.SliceStable(transactions, func(i, j int) bool {
		return transactions[i].ID < transactions[j].ID
	})

	var transactionList []transactiondto.TransactionResponse
	currentID := 0
	var transactionToConvert [][]models.TransactionRawQuery
	var listRawTransactionByID []models.TransactionRawQuery

	for idx, transaction := range transactions {
		if currentID == transaction.ID {
			listRawTransactionByID = append(listRawTransactionByID, transaction)
		} else {
			if len(listRawTransactionByID) > 0 {
				transactionToConvert = append(transactionToConvert, listRawTransactionByID)
				listRawTransactionByID = []models.TransactionRawQuery{}
			}
			listRawTransactionByID = append(listRawTransactionByID, transaction)
		}
		if idx+1 == len(transactions) {
			transactionToConvert = append(transactionToConvert, listRawTransactionByID)
		}
		currentID = transaction.ID
	}
	for _, v := range transactionToConvert {
		transactionList = append(transactionList, convertTransactionRawToResponse(v))
	}
	return transactionList
}

func convertTransactionRawToResponse(t []models.TransactionRawQuery) transactiondto.TransactionResponse {
	var result = transactiondto.TransactionResponse{
		ID: t[0].ID,
		User: transactiondto.UserTransactionResponse{
			ID:       t[0].UserID,
			Fullname: t[0].Fullname,
			Email:    t[0].Email,
		},
		UserID:    t[0].UserID,
		Status:    t[0].Status,
		CreatedAt: t[0].CreatedAt,
		UpdatedAt: t[0].UpdatedAt,
	}
	var products []cartdto.CartResponse
	for _, transaction := range t {
		cart := cartdto.CartResponse{
			ID:            transaction.ProductID,
			OrderQuantity: transaction.OrderQuantity,
			Name:          transaction.ProductName,
			Price:         transaction.Price,
			Description:   transaction.Description,
			Photo:         transaction.Photo,
		}
		products = append(products, cart)
	}
	result.Cart = products
	return result

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
