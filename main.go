package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", welcome)
	e.Logger.Fatal(e.Start(":1323"))
}

func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Echo!")
}
