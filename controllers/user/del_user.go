package user

import (
	"net/http"
	"prakerja12/configs"
	basemodel "prakerja12/models/base"
	usermodel "prakerja12/models/user"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DeleteController(c echo.Context) error {
	var users []usermodel.User
	result := configs.DB.Delete(&users)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Handle case where no records were found to delete
			return c.JSON(http.StatusNotFound, basemodel.BaseResponse{
				Status:  false,
				Message: "No data found to delete",
				Data:    nil,
			})
		}

		// Handle other delete errors
		return c.JSON(http.StatusInternalServerError, basemodel.BaseResponse{
			Status:  false,
			Message: "Failed to delete data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, basemodel.BaseResponse{
		Status:  true,
		Message: "Success",
		Data:    users,
	})
}
