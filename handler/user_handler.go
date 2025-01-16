package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "Gia Bao",
		"email": "hpgbao@gmail.com",
	})
}

func HandleSignUp(c echo.Context) error {
	type User struct {
		Email    string `json:"email"`
		FullName string `json:"fullName"`
		Name     string `json:"name"`
	}

	user := User{
		Email:    "hpgbao@gmail.com",
		FullName: "Gia Bao",
		Name:     "Gia Bao",
	}

	return c.JSON(http.StatusOK, user)
}
