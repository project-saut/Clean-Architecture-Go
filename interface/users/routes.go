package users

import (
	database "Belajar-Go-Echo/infrastructure/database/sqlite"
	m "Belajar-Go-Echo/infrastructure/http/middleware"
	ur "Belajar-Go-Echo/repository/users"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {
	db := database.ConnectDB()
	userRepo := ur.NewUserRepository(db)
	userHandler := NewUserHandler(userRepo)

	echo.GET("/users", userHandler.GetAllUserHandler, m.JWTMiddleware())
	echo.POST("/users", userHandler.CreateUserHandler)
}
