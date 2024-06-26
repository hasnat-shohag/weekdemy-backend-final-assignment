package controllers

import (
	"net/http"
	"noob-server/pkg/domain"
	"noob-server/pkg/types"

	"github.com/labstack/echo/v4"
)

// IAuthController is an interface that defines the methods implemented by the AuthController struct.
type IAuthController interface {
	Login(e echo.Context) error
	Signup(e echo.Context) error
}

// AuthController defines the methods of the IAuthController interface.
type AuthController struct {
	authSvc domain.IAuthService
}

// NewAuthController is a function that returns a new instance of the AuthController struct.
func NewAuthController(authSvc domain.IAuthService) AuthController {
	return AuthController{
		authSvc: authSvc,
	}
}

func (authController *AuthController) Login(e echo.Context) error {
	// bind the request body to the LoginRequest struct
	loginRequest := &types.LoginRequest{}
	if err := e.Bind(loginRequest); err != nil {
		return e.JSON(http.StatusBadRequest, "invalid request body")
	}

	// validate the request body
	if err := loginRequest.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	// pass the request to the service layer
	loginResponse, err := authController.authSvc.LoginUser(loginRequest)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, loginResponse)
}

func (authController *AuthController) Signup(e echo.Context) error {
	// bind the request body to the SignupRequest struct
	registerRequest := &types.SignupRequest{}
	if err := e.Bind(registerRequest); err != nil {
		return e.JSON(http.StatusBadRequest, "invalid request body")
	}

	// validate the request body
	if err := registerRequest.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	// pass the request to the service layer
	if err := authController.authSvc.SignupUser(registerRequest); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, "user was created successfully")
}
