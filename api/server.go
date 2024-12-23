package api

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/timothypattikawa/amole-services/cart-service/internal/config"
	"github.com/timothypattikawa/amole-services/cart-service/internal/handler"
	"github.com/timothypattikawa/amole-services/cart-service/pkg/errors"
)

func RunServer(h *handler.CartHandler, serverConfig config.ServerConfig) {
	e := echo.New()

	e.HideBanner = true
	e.HTTPErrorHandler = errors.CostumeError

	handler.Handler(e, h)

	err := e.Start(":" + serverConfig.Server)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

}
