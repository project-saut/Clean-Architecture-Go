package users

import (
	model "Belajar-Go-Echo/domains/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	repository model.UserRepository
}

func NewUserHandler(repo model.UserRepository) model.UserHandler {
	return &handler{
		repository: repo,
	}
}

func (h *handler) CreateUserHandler(c echo.Context) error {
	user := model.User{}

	e := c.Bind(&user)
	if e != nil {
		return e
	}

	err := h.repository.CreateUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create user",
		"Users":   user,
	})
}

func (h *handler) GetAllUserHandler(c echo.Context) error {
	users := h.repository.GetAllUser()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success get all user",
		"Users":   users,
	})
}

func (h *handler) GetUserByIdHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user, err := h.repository.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success get user by id",
		"User":    user,
	})
}
