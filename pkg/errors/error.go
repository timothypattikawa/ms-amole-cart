package errors

import (
	"errors"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/timothypattikawa/amole-services/cart-service/internal/dto"
)

func CostumeError(err error, c echo.Context) {
	var e *BusinessProcessError
	ok := errors.As(err, &e)

	if ok {
		err := c.JSON(e.status, dto.BaseResponse{
			Status: e.status,
			Error:  e.Error(),
		})

		if err != nil {
			log.Println("error json response for exception")
		}
	}
}
