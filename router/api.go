package router

import (
	"backend-github-trending/handler"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	api.Echo.GET("/users/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/users/sign-up", api.UserHandler.HandleSignUp)
}
