package auth

import (
	"net/http"
	"prakerja12/configs"
	"prakerja12/middleware"
	"prakerja12/models/base"
	usermodel "prakerja12/models/user"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(c echo.Context) error {
	var user usermodel.User
	c.Bind(&user)

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hashPassword)

	result := configs.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed add data to database",
			Data:    nil,
		})
	}

	var authResponse = usermodel.ResponseAuth{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Token: middleware.GenerateTokenJWT(user.Id, user.Name),
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success register",
		Data:    authResponse,
	})
}
