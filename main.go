package main

import (
	"backend-github-trending/db"
	"backend-github-trending/handler"
	"github.com/labstack/echo/v4"
)

func main() {

	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "hpgbao",
		Password: "Baolodc2204.",
		DBName:   "postgres",
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.GET("/", handler.Welcome)

	e.GET("/users/sign-in", handler.HandleSignIn)
	e.GET("/users/sign-up", handler.HandleSignUp)

	e.Logger.Fatal(e.Start(":3000"))
}
