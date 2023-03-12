package handlers

import (
	dtoaddress "backendwaysbeans/dto/address"
	dto "backendwaysbeans/dto/result"
	"backendwaysbeans/models"
	"backendwaysbeans/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerAddress struct {
	AddressRepository repositories.AddressRepository
}

func HandlerAddress(AddressRepository repositories.AddressRepository) *handlerAddress {
	return &handlerAddress{
		AddressRepository: AddressRepository,
	}
}

func (h *handlerAddress) FindAddress(c echo.Context) error {
	addresses, err := h.AddressRepository.FindAddress()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: addresses})
}

func (h *handlerAddress) FindAddressByUser(c echo.Context) error {
	userLogin := c.Get("userLogin")

	user_id := userLogin.(jwt.MapClaims)["id"].(float64)
	addresses, err := h.AddressRepository.FindAddressByUser(int(user_id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: addresses})
}

func (h *handlerAddress) GetAddress(c echo.Context) error {
	addressID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	address, err := h.AddressRepository.GetAddress(addressID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: address})
}

func (h *handlerAddress) CreateAddressByUser(c echo.Context) error {
	userLogin := c.Get("userLogin")

	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	request := new(dtoaddress.CreateAddressRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	newAddress := models.Addresses{
		Name:      request.Name,
		Phone:     request.Phone,
		Address:   request.Address,
		PostCode:  request.PostCode,
		UserID:    int(userId),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	address, err := h.AddressRepository.CreateAddress(newAddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: address})
}

func (h *handlerAddress) DeleteAddress(c echo.Context) error {
	addressID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	address, err := h.AddressRepository.GetAddress(addressID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	data, err := h.AddressRepository.DeleteAddress(address, addressID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: data})
}

func (h *handlerAddress) UpdateAddress(c echo.Context) error {
	addressID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	request := new(dtoaddress.UpdateAddressRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err = validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	address, err := h.AddressRepository.GetAddress(addressID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	updatedAddress := models.Addresses{
		ID:        address.ID,
		Name:      request.Name,
		Phone:     request.Phone,
		Address:   request.Address,
		PostCode:  request.PostCode,
		UserID:    address.UserID,
		CreatedAt: address.CreatedAt,
		UpdatedAt: time.Now(),
	}

	newAddress, err := h.AddressRepository.UpdateAddress(updatedAddress)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: newAddress})
}
