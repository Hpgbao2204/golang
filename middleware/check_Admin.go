package middleware

import (
	"backend-github-trending/model"
	"backend-github-trending/model/req"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Handle logic in here
			req := req.SignIn{}
			var err error
			if err = c.Bind(&req); err != nil {
				log.Error(err.Error())
				return c.JSON(http.StatusBadRequest, model.Respone{
					StatusCode: http.StatusBadRequest,
					Message:    err.Error(),
					Data:       nil,
				})
			}
			if req.Email != "admin@gmail.com" {
				return c.JSON(http.StatusUnauthorized, model.Respone{
					StatusCode: http.StatusUnauthorized,
					Message:    "ban k phai la admin",
					Data:       nil,
				})
			}
			return next(c)
		}
	}
}
