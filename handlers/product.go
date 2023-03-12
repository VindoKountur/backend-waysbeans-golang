package handlers

import (
	productsdto "backendwaysbeans/dto/products"
	dto "backendwaysbeans/dto/result"
	"backendwaysbeans/models"
	"backendwaysbeans/repositories"
	"net/http"
	"time"

	"strconv"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var path_file = "/uploads/"

type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

func (h *handlerProduct) FindProducts(c echo.Context) error {
	searchName := c.QueryParam("name")
	var products []models.Product
	var err error
	if searchName == "" {
		products, err = h.ProductRepository.FindProduct()
	} else {
		products, err = h.ProductRepository.FindProductByName(searchName)
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: products})
}

func (h *handlerProduct) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: product})
}

func (h *handlerProduct) CreateProduct(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)

	userLogin := c.Get("userLogin")

	idUserLogin := (userLogin.(jwt.MapClaims)["id"]).(float64)

	price, _ := strconv.Atoi(c.FormValue("price"))
	stock, _ := strconv.Atoi(c.FormValue("stock"))

	request := productsdto.CreateProductRequest{
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
		Price:       price,
		Stock:       stock,
		UserID:      int(idUserLogin),
		Photo:       dataFile,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	product := models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Photo:       request.Photo,
		Stock:       request.Stock,
		UserID:      request.UserID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	product, err = h.ProductRepository.CreateProduct(product)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertResponseProduct(product)})
}

func (h *handlerProduct) UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	dataFile := c.Get("dataFile").(string)

	price, _ := strconv.Atoi(c.FormValue("price"))
	stock, _ := strconv.Atoi(c.FormValue("stock"))

	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	request := productsdto.UpdateProductRequest{
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
		Price:       price,
		Stock:       stock,
		Photo:       dataFile,
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	if request.Name != "" {
		product.Name = request.Name
	}
	if request.Description != "" {
		product.Description = request.Description
	}
	if request.Price != 0 {
		product.Price = request.Price
	}
	if request.Stock != 0 {
		product.Stock = request.Stock
	}
	if request.Photo != "" {
		product.Photo = request.Photo
	}

	product.UpdatedAt = time.Now()

	product, err = h.ProductRepository.UpdateProduct(product)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertResponseProduct(product)})
}

func (h *handlerProduct) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetProduct(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	product, err = h.ProductRepository.DeleteProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	// Delete Photo
	// err = os.Remove(path_file + product.Photo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertResponseProduct(product)})
}

func convertResponseProduct(p models.Product) productsdto.ProductResponse {
	return productsdto.ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Photo:       p.Photo,
		Stock:       p.Stock,
	}
}
