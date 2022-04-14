package users

import "github.com/labstack/echo/v4"

type UserHandler interface {
	CreateUserHandler(c echo.Context) error
	GetAllUserHandler(c echo.Context) error
	GetUserByIdHandler(c echo.Context) error
}
