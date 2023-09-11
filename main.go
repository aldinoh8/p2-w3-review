package main

import (
	"example/config"
	"example/controller"
	"example/middleware"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := config.InitDb(os.Getenv("DB"))

	auth := middleware.Auth{DB: db}

	user := e.Group("/users")
	userController := controller.User{DB: db}
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
	}

	player := e.Group("/players")
	playerController := controller.Player{DB: db}
	{
		player.GET("", playerController.Index)
		player.POST("", auth.Authenticate(playerController.Create))
		player.PUT("/:id", auth.Authenticate(playerController.Update))
	}

	e.Logger.Fatal(e.Start(":8000"))
}
