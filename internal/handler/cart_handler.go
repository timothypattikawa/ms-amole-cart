package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/timothypattikawa/amole-services/cart-service/internal/dto"
	"github.com/timothypattikawa/amole-services/cart-service/internal/service"
	exception "github.com/timothypattikawa/amole-services/cart-service/pkg/errors"
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
	atcRequest := new(dto.AddToCartRequest)
	err := e.Bind(atcRequest)
	if err != nil {
		log.Printf("fail bind atc cause err {%v}", err)
		return exception.NewBusinessProcessError("Somthing wen't wrong", http.StatusInternalServerError)
	}

	responseATC, err := ch.cr.AddToCart(e.Request().Context(), *atcRequest)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, dto.BaseResponse{
		Status: http.StatusOK,
		Data:   responseATC,
	})
}

func (ch CartHandler) GetAllListCart(e echo.Context) error {
	return e.JSON(http.StatusOK, dto.BaseResponse{
		Data: nil,
	})
}
