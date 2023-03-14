package handlers

import (
	cartdto "backendwaysbeans/dto/cart"
	dto "backendwaysbeans/dto/result"
	transactiondto "backendwaysbeans/dto/transaction"
	"backendwaysbeans/models"
	"backendwaysbeans/repositories"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gopkg.in/gomail.v2"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
	CartRepository        repositories.CartRepository
	ProductRepository     repositories.ProductRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository, CartRepository repositories.CartRepository, ProductRepository repositories.ProductRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository, CartRepository, ProductRepository}
}

func SendMail(status string, transaction models.Transaction) {

	if status != transaction.Status && (status == "success") {
		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "Waysbeans <demo.dumbways@gmail.com>"
		var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
		var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

		var listProduct []string
		var totalPrice int
		for _, cart := range transaction.Cart {
			productName := cart.Product.Name
			price := strconv.Itoa(cart.Product.Price)
			qty := strconv.Itoa(cart.OrderQuantity)
			subtotal := strconv.Itoa(cart.OrderQuantity * cart.Product.Price)
			totalPrice += cart.Product.Price * cart.OrderQuantity

			var td = fmt.Sprintf(`
			<tr>
				<td style='padding: 3px;'>%s</td>
				<td style='padding: 3px;'>%s</td>
				<td style='padding: 3px;'>%s</td>
				<td style='padding: 3px;'>%s</td>
			</tr>`, productName, price, qty, subtotal)
			listProduct = append(listProduct, td)
		}

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", transaction.User.Email)
		mailer.SetHeader("Subject", "Transaction Status")
		mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
	  <html lang="en">
		<head>
		<meta charset="UTF-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Waysbeans Transaction</title>
		<style>
		  h1 {
		  color: brown;
		  }
		</style>
		</head>
		<body>
		<h2>Product payment Success</h2>
		<table border="0" cellspacing="0" style="margin-top: 10px">
		<thead>
			<tr>
                <th style='padding: 3px;'>Product Name</th>
                <th style='padding: 3px;'>Price</th>
                <th style='padding: 3px;'>Quantity</th>
                <th style='padding: 3px;'>Subtotal</th>
            </tr>
        </thead>
		<tbody>
			%s
		</tbody>
		</table>
		<h2>Total Payment : %s</h2>
		</body>
	  </html>`, strings.Join(listProduct, ""), strconv.Itoa(totalPrice)))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent! to " + transaction.User.Email)
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

	// Generate ID
	var transactionIsMatch = false
	var transactionId int
	for !transactionIsMatch {
		transactionId = int(time.Now().Unix())
		transactionData, _ := h.TransactionRepository.GetTransaction(transactionId)
		if transactionData.ID == 0 {
			transactionIsMatch = true
		}
	}

	totalPrice := 0
	// Handle Products
	for _, product := range request.Products {
		getProduct, err := h.ProductRepository.GetProduct(product.ProductID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		}

		if getProduct.Stock == 0 {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: fmt.Sprintf("Product %s is not in stock", getProduct.Name)})
		}

		if getProduct.Stock < product.OrderQuantity {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: fmt.Sprintf("Failed to make transaction, not enough stock for product : %s", getProduct.Name)})
		}
		totalPrice += getProduct.Price * product.OrderQuantity
	}

	newTransactions := models.Transaction{
		ID:         transactionId,
		UserID:     int(idUserLogin),
		Name:       request.Name,
		Email:      request.Email,
		Address:    request.Address,
		PostCode:   request.PostCode,
		Phone:      request.Phone,
		TotalPrice: totalPrice,
		Status:     "waiting",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	createdTransaction, err := h.TransactionRepository.CreateTransaction(newTransactions)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	for _, product := range request.Products {
		newCart := models.Cart{
			TransactionID: transactionId,
			ProductID:     product.ProductID,
			OrderQuantity: product.OrderQuantity,
		}
		_, err := h.CartRepository.CreateCart(newCart)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		}
	}

	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(createdTransaction.ID),
			GrossAmt: int64(createdTransaction.TotalPrice),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: createdTransaction.User.Name,
			Email: createdTransaction.User.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: snapResp})

	// data, err := h.TransactionRepository.GetTransaction(createdTransaction.ID)
	// return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: (data)})
}

func (h *handlerTransaction) FindTransaction(c echo.Context) error {
	transactions, err := h.TransactionRepository.FindTransaction()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	var response []transactiondto.TransactionResponse
	for _, transaction := range transactions {
		response = append(response, convertTransactionModelToResponse(transaction))
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: response})
}

func (h *handlerTransaction) GetTransaction(c echo.Context) error {
	transactionID, _ := strconv.Atoi(c.Param("id"))
	transactions, err := h.TransactionRepository.GetTransaction(transactionID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertTransactionModelToResponse(transactions)})
}

func (h *handlerTransaction) GetUserTransaction(c echo.Context) error {
	idUserLogin := int((c.Get("userLogin").(jwt.MapClaims)["id"]).(float64))
	transactions, err := h.TransactionRepository.FindTransactionByUserID(idUserLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	var response []transactiondto.TransactionResponse
	for _, transaction := range transactions {
		response = append(response, convertTransactionModelToResponse(transaction))
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: response})
}

func (h *handlerTransaction) Notification(c echo.Context) error {
	var notificationPayload map[string]interface{}

	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	order_id, _ := strconv.Atoi(orderId)

	fmt.Print("ini payloadnya", notificationPayload)
	transaction, _ := h.TransactionRepository.GetTransaction(order_id)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			h.TransactionRepository.UpdateTransaction("pending", order_id)
		} else if fraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			SendMail("success", transaction)
			h.TransactionRepository.UpdateTransaction("success", order_id)
		}
	} else if transactionStatus == "settlement" {
		// TODO set transaction status on your databaase to 'success'
		SendMail("success", transaction)
		h.TransactionRepository.UpdateTransaction("success", order_id)
	} else if transactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
		h.TransactionRepository.UpdateTransaction("failed", order_id)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		// TODO set transaction status on your databaase to 'failure'
		h.TransactionRepository.UpdateTransaction("failed", order_id)
	} else if transactionStatus == "pending" {
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		h.TransactionRepository.UpdateTransaction("pending", order_id)
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Status", Data: notificationPayload})
}

func convertCartResponse(c []models.Cart) []cartdto.CartResponse {
	var cartResponse []cartdto.CartResponse
	for _, cart := range c {
		cartResponse = append(cartResponse, cartdto.CartResponse{
			ID:            cart.Product.ID,
			Name:          cart.Product.Name,
			Price:         cart.Product.Price,
			Description:   cart.Product.Description,
			Photo:         cart.Product.Photo,
			OrderQuantity: cart.OrderQuantity,
		})
	}
	return cartResponse
}

func convertTransactionModelToResponse(t models.Transaction) transactiondto.TransactionResponse {
	return transactiondto.TransactionResponse{
		ID: t.ID,
		User: transactiondto.UserTransactionResponse{
			ID:       t.User.ID,
			Fullname: t.User.Name,
			Email:    t.User.Email,
		},
		Name:      t.Name,
		Email:     t.Email,
		Address:   t.Address,
		Phone:     t.Phone,
		PostCode:  t.PostCode,
		Status:    t.Status,
		CreatedAt: t.CreatedAt,
		Cart:      convertCartResponse(t.Cart),
	}
}
