package main

import (
	"app-fullstack-gotth/handlers/authhandler"
	"app-fullstack-gotth/share"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// files static
	e.Static("/static", "static")
	e.HTTPErrorHandler = share.CustomHTTPErrorHandler

	// routes
	e.GET("/login", authhandler.NewGetLoginHandler().Serve)
	e.GET("/register", authhandler.NewGetRegisterHandler().Serve)

	// Start Server
	e.Logger.Fatal(e.Start(":8082"))
}
