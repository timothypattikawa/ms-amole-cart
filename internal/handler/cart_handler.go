package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/timothypattikawa/amole-services/cart-service/internal/dto"
	"github.com/timothypattikawa/amole-services/cart-service/internal/service"
)

type CartHandler struct {
	cr service.CartService
}

func NewCartHandler(cr service.CartService) *CartHandler {
	return &CartHandler{
		cr: cr,
	}
}

func (ch CartHandler) AddTocart(e echo.Context) error {

	return e.JSON(http.StatusOK, dto.BaseResponse{
		Data: "succes add to cart",
	})
}

func (ch CartHandler) GetAllListCart(e echo.Context) error {
	return e.JSON(http.StatusOK, dto.BaseResponse{
		Data: nil,
	})
}
