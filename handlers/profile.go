package handlers

import (
	profiledto "backendwaysbeans/dto/profile"
	dto "backendwaysbeans/dto/result"
	"backendwaysbeans/models"
	"backendwaysbeans/repositories"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerProfile struct {
	ProfileRepository repositories.ProfileRepository
}

func HandlerProfile(ProfileRepository repositories.ProfileRepository) *handlerProfile {
	return &handlerProfile{ProfileRepository}
}

func (h *handlerProfile) UpdateProfileByUser(c echo.Context) error {
	userLogin := c.Get("userLogin")
	idUserLogin := userLogin.(jwt.MapClaims)["id"].(float64)
	dataFilePhoto := c.Get("dataFile").(string)

	request := profiledto.ProfileUpdateRequest{
		Photo: dataFilePhoto,
		Phone: c.FormValue("phone"),
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	profile := models.Profile{
		UserID:    int(idUserLogin),
		Photo:     request.Photo,
		Phone:     request.Phone,
		UpdatedAt: time.Now(),
	}

	updatedProfile, err := h.ProfileRepository.UpdateProfileByUser(profile, int(idUserLogin))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: updatedProfile})
}
