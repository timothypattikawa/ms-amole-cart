package handler

import "github.com/labstack/echo/v4"

func Handler(e *echo.Echo, h *CartHandler) {
	e.POST("/v1/add-to-cart", h.AddTocart)
	e.GET("/v1/cart", h.GetAllListCart)
}
