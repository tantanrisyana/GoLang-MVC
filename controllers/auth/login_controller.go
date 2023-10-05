package auth

import (
	"net/http"
	"prakerja12/models/base"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success register",
		Data:    nil,
	})
}
