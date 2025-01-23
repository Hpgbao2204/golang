package handler

import (
	"backend-github-trending/model"
	req "backend-github-trending/model/req"
	"backend-github-trending/repository"
	security "backend-github-trending/security"
	validator "github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	_ "golang.org/x/mod/modfile"
	"net/http"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

func (u *UserHandler) HandleSignUp(c echo.Context) error {
	req := req.SignUp{}

	var err error

	if err = c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Respone{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Respone{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	hash := security.HashAndSalt([]byte(req.Password))
	role := model.MEMBER.String()

	userId, err := uuid.NewUUID()
	if err != nil {
		return c.JSON(http.StatusForbidden, model.Respone{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user := model.User{
		UserID:   userId.String(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: hash,
		Role:     role,
		Token:    "",
	}

	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Respone{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user.Password = ""

	return c.JSON(http.StatusOK, model.Respone{
		StatusCode: http.StatusOK,
		Message:    "Sign up successfully",
		Data:       user,
	})
}

func (u *UserHandler) HandleSignIn(c echo.Context) error {
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

	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Respone{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user, err := u.UserRepo.CheckLogin(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Respone{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// compare 2 pass
	isTheSame := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isTheSame {
		return c.JSON(http.StatusUnauthorized, model.Respone{
			StatusCode: http.StatusUnauthorized,
			Message:    "Dang nhap that bai vi sai mk",
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Respone{
		StatusCode: http.StatusOK,
		Message:    "Sign In successfully",
		Data:       nil,
	})
}
