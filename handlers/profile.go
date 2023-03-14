package handlers

import (
	profiledto "backendwaysbeans/dto/profile"
	dto "backendwaysbeans/dto/result"
	"backendwaysbeans/repositories"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
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

	profile, err := h.ProfileRepository.GetProfileByUser(int(idUserLogin))

	request := profiledto.ProfileUpdateRequest{
		Photo: dataFilePhoto,
		Phone: c.FormValue("phone"),
	}
	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	if request.Phone != "" {
		profile.Phone = request.Phone
	}
	if dataFilePhoto != "" {
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctx, request.Photo, uploader.UploadParams{Folder: "waysbeans-profile"})

		if err != nil {
			fmt.Println(err.Error())
		}
		profile.Photo = resp.SecureURL
	}
	profile.UpdatedAt = time.Now()

	updatedProfile, err := h.ProfileRepository.UpdateProfileByUser(profile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: updatedProfile})
}
