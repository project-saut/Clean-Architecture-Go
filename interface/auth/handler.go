package auth

import (
	"Belajar-Go-Echo/domains/auth"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type handler struct {
	repostory auth.AuthRepository
}

func NewAuthHandler(repo auth.AuthRepository) auth.AuthHandler {
	return &handler{
		repostory: repo,
	}
}

func (h *handler) LoginHandler(c echo.Context) error {
	credential := auth.Auth{}

	err := c.Bind(&credential)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.repostory.Login(credential)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, e := token.SignedString([]byte("saut"))

	if e != nil {
		return c.JSON(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Create JWT",
		"Token":   jwtToken,
	})
}
