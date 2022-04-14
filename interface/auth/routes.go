package auth

import (
	database "Belajar-Go-Echo/infrastructure/database/sqlite"
	ar "Belajar-Go-Echo/repository/auth"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {
	db := database.ConnectDB()
	authRepo := ar.NewAuthRepository(db)
	authHandler := NewAuthHandler(authRepo)

	echo.POST("/auth/login", authHandler.LoginHandler)
}
